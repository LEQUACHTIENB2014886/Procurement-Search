<template>
  <div class="container">
    <!-- Hiển thị Loading nếu đang tìm kiếm -->
    <Loading v-if="isLoading" />

    <el-row class="header">
      <el-col :span="4"></el-col>
      <el-col :span="15" class="flex-container">
        <div class="input-container">
          <h2>BUYNO</h2>
          <el-input
            v-model="input"
            placeholder="202407, press Enter"
            clearable
            @keyup.enter="fetchProcurementData"
            @clear="resetSearch"
            class="input"
          />
        </div>
        <div class="select-container">
          <h2>GSBH</h2>
          <el-select
            v-model="gsbhValue"
            class="select"
            @change="fetchProcurementData"
          >
            <el-option label="VA12" value="VA12"></el-option>
            <el-option label="VC02" value="VC02"></el-option>
          </el-select>
        </div>
        <div class="average-container">
          <h2>
            DIFF DAY AVERAGE:
            <span class="average-value">{{ formattedAverageDiffDay }}</span>
          </h2>
        </div>
      </el-col>
      <el-col :span="5"></el-col>
    </el-row>

    <div class="table-container">
      <el-auto-resizer>
        <template #default="{ width, height }">
          <el-table-v2
            :columns="columns"
            :data="filteredData"
            :width="width > 1200 ? 1200 : width"
            :height="height"
            fixed
            row-key="stt"
            class="custom-table"
          >
            <el-table-column
              v-for="column in columns"
              :key="column.key"
              :prop="column.dataKey"
              :label="column.title"
              :width="column.width"
            />
          </el-table-v2>
        </template>
      </el-auto-resizer>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed } from "vue";
import axios from "axios";
import Loading from "./Loading.vue"; // Import Loading component

interface ProcurementItem {
  STT: string;
  DDBH: string;
  StartDate: string;
  EndDate: string;
  DiffDay: string;
}

const input = ref<string>(""); // Giá trị input
const gsbhValue = ref<string>("VA12"); // Giá trị GSBH
const procurementData = ref<ProcurementItem[]>([]); // Dữ liệu lấy từ API
const filteredData = ref<ProcurementItem[]>([]); // Dữ liệu đã lọc theo điều kiện
const averageDiffDay = ref<number>(0); // Trung bình ngày chênh lệch
const isLoading = ref<boolean>(false); // Biến trạng thái Loading

const columns = [
  { key: "stt", dataKey: "STT", title: "STT", width: 150 },
  { key: "ddbh", dataKey: "DDBH", title: "DDBH", width: 250 },
  { key: "startDate", dataKey: "StartDate", title: "Start Date", width: 300 },
  { key: "endDate", dataKey: "EndDate", title: "End Date", width: 300 },
  { key: "diffDay", dataKey: "DiffDay", title: "Diff Day", width: 170 },
];

// Lấy dữ liệu từ API

// Cập nhật thời gian chờ trong Table.vue
const fetchProcurementData = async () => {
  if (!input.value) return; // Không gọi API nếu input rỗng

  isLoading.value = true; // Bắt đầu hiển thị Loading

  try {
    const response = await axios.get(
      `http://localhost:8081/api/v1/pro/procurement`,
      {
        params: {
          BUYNO: input.value,
          GSBH: gsbhValue.value,
        },
      }
    );

    const data = response.data.data || [];
    procurementData.value = data as ProcurementItem[];
    filteredData.value = data.map((item: ProcurementItem, index: number) => ({
      ...item,
      STT: (index + 1).toString(),
    }));

    calculateAverageDiffDay();
  } catch (error) {
    console.error("Error fetching data:", error);
  } finally {
    // Tăng thời gian loading thêm 0.5s sau khi có dữ liệu
    setTimeout(() => {
      isLoading.value = false; // Sau 0.5s, ẩn loading
    }, 500);
  }
};



// Tính trung bình DiffDay
const calculateAverageDiffDay = () => {
  if (!filteredData.value.length) {
    averageDiffDay.value = 0;
    return;
  }
  const totalDays = filteredData.value.reduce(
    (acc, item) => acc + Number(item.DiffDay || 0),
    0
  );
  averageDiffDay.value = parseFloat(
    (totalDays / filteredData.value.length).toFixed(2)
  );
};

const formattedAverageDiffDay = computed(() => {
  return averageDiffDay.value < 10
    ? `0${averageDiffDay.value.toFixed(2)}`
    : averageDiffDay.value.toFixed(2);
});

// Reset tìm kiếm
const resetSearch = () => {
  input.value = "";
  gsbhValue.value = "VA12";
  filteredData.value = procurementData.value.map((item, index) => ({
    ...item,
    STT: (index + 1).toString(),
  }));
  calculateAverageDiffDay();
};

onMounted(() => {
  fetchProcurementData();
});
</script>

<style scoped>
.container {
  height: 89vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

.header {
  width: 100%;
  height: 140px;
  margin-top: -40px;
}

.flex-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
}

.input-container,
.select-container,
.average-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

.average-value {
  color: red;
  font-weight: bold;
}

.input,
.select {
  font-size: 18px;
}

.input {
  width: 250px;
}

.select {
  width: 80px;
}

.table-container {
  flex: 1;
  width: 100%;
  max-width: 1200px; 
  display: flex;
  justify-content: center;
  overflow: hidden;
}

.el-table {
  width: 100%;
  table-layout: fixed; 
}

.el-table th,
.el-table td {
  font-size: 22px; /* Tăng kích thước chữ trong bảng */
}

.el-table-column {
  padding: 15px; /* Tăng padding để cột rộng hơn */
}

.el-table th {
  text-align: center;
}

.el-table td {
  text-align: center;
}
</style>
