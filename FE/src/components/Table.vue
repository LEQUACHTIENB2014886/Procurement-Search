<template>
  <div class="container">
    <Loading v-if="isLoading" />
    <el-row class="header">
      <el-col :span="4" />
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
            <el-option label="VA12" value="VA12" />
            <el-option label="VC02" value="VC02" />
          </el-select>
        </div>
        <div class="average-container">
          <h2>
            DIFF DAY AVERAGE:
            <span class="average-value">{{ formattedAverageDiffDay }}</span>
          </h2>
        </div>
      </el-col>
      <el-col :span="5" />
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

<script setup>
import { ref, onMounted, computed } from "vue";
import axios from "axios";
import Loading from "./Loading.vue";


const input = ref("");
const gsbhValue = ref("VA12"); 
const procurementData = ref([]); 
const filteredData = ref([]); 
const averageDiffDay = ref(0); 
const isLoading = ref(false); 

// Định nghĩa cấu trúc cột bảng
const columns = [
  { key: "stt", dataKey: "STT", title: "STT", width: 150 },
  { key: "ddbh", dataKey: "DDBH", title: "DDBH", width: 250 },
  { key: "startDate", dataKey: "StartDate", title: "Start Date", width: 300 },
  { key: "endDate", dataKey: "EndDate", title: "End Date", width: 300 },
  { key: "diffDay", dataKey: "DiffDay", title: "Diff Day", width: 170 },
];

// Hàm gọi API để lấy dữ liệu procurement và xử lý dữ liệu trả về
const fetchProcurementData = async () => {
  if (!input.value) return; 
  isLoading.value = true; 

  try {
    // Gửi yêu cầu GET tới API để lấy dữ liệu
    const response = await axios.get(
      `http://localhost:8081/api/v1/pro/procurement`, // Địa chỉ API
      // `http://192.168.40.105:8081/api/v1/procurement/get`,
      {
        params: {
          BUYNO: input.value, 
          GSBH: gsbhValue.value, 
        },
      }
    );

    // Lấy dữ liệu từ API và lưu vào p rocurementData
    const data = response.data.data || [];
    procurementData.value = data;
    filteredData.value = data.map((item, index) => ({
      ...item,
      STT: (index + 1).toString(), 
    }));

    calculateAverageDiffDay(); // Tính toán giá trị trung bình DiffDay
  } catch (error) {
    console.error("Error fetching data:", error); // Xử lý lỗi nếu có
  } finally {
    isLoading.value = false; // Kết thúc quá trình tải dữ liệu
  }
};

// Hàm tính toán giá trị trung bình DiffDay từ dữ liệu filteredData
const calculateAverageDiffDay = () => {
  if (!filteredData.value.length) {
    averageDiffDay.value = 0; // Nếu không có dữ liệu, đặt giá trị trung bình là 0
    return;
  }
  const totalDays = filteredData.value.reduce(
    (acc, item) => acc + Number(item.DiffDay || 0), // Tính tổng số ngày chênh lệch
    0
  );
  averageDiffDay.value = parseFloat(
    (totalDays / filteredData.value.length).toFixed(2) // Tính giá trị trung bình và làm tròn đến 2 chữ số
  );
};

// Computed property để định dạng giá trị trung bình DiffDay
const formattedAverageDiffDay = computed(() => {
  return averageDiffDay.value < 10
    ? `0${averageDiffDay.value.toFixed(2)}` 
    : averageDiffDay.value.toFixed(2);
});

// Hàm reset lại giá trị tìm kiếm và bảng dữ liệu
const resetSearch = () => {
  input.value = ""; 
  gsbhValue.value = "VA12"; 
  filteredData.value = procurementData.value.map((item, index) => ({
    ...item,
    STT: (index + 1).toString(), 
  }));
  calculateAverageDiffDay(); 
};

// Lifecycle hook để lấy dữ liệu khi component được mount
onMounted(() => {
  fetchProcurementData(); 
});
</script>

<style scoped>
.container {
  height: 90vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 20px;
  margin-top:-20px;
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

.input-container{
  margin-left:20px;
}

.select-container{
  margin-left:-450px;
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
  width: 160px;
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
  margin-top:-30px;
}

.el-table {
  width: 100%;
  table-layout: fixed;
}

.el-table th,
.el-table td {
  font-size: 22px;
}

.el-table-column {
  padding: 15px;
}

.el-table th {
  text-align: center;
}

.el-table td {
  text-align: center;
}

</style>
