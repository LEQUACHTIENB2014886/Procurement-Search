<template>
    <div>
      <el-row>
        <el-col :span="3"></el-col>
        <el-col :span="6" class="flex-start">
          <h2>BUYNO</h2>
          <el-input
            v-model="input"
            placeholder="Nhập BUYNO..."
            clearable
            @keyup.enter="searchData"  
            @clear="resetSearch" 
            class="input"
            style="max-width: 200px"
          />
        </el-col>
        <el-col :span="5" class="flex-start">
          <h2>GSBH</h2>
          <el-select
            v-model="gsbhValue"
            class="input"
            style="max-width: 90px"
            @change="searchData"
          >
            <el-option label="Tất cả" value="Tất cả"></el-option>
            <el-option label="VA12" value="VA12"></el-option>
            <el-option label="VC02" value="VC02"></el-option>
          </el-select>
        </el-col>
        <el-col :span="7" class="flex-end">
          <h2>
            DIFF DAY AVERAGE: <a style="color: red">{{ averageDiffDay }}</a>
          </h2>
        </el-col>
        <el-col :span="3"></el-col>
      </el-row>
  
      <el-row class="table">
        <el-col :span="3"></el-col>
        <el-col :span="2" class="table-cell bold-text bg-dark"><a>STT</a></el-col>
        <el-col :span="2" class="table-cell bold-text bg-dark"><a>BUYNO</a></el-col>
        <el-col :span="3" class="table-cell bold-text bg-dark"><a>GSBH</a></el-col>
        <el-col :span="3" class="table-cell bold-text bg-dark"><a>DDBH</a></el-col>
        <el-col :span="3" class="table-cell bold-text bg-dark"><a>Start Date</a></el-col>
        <el-col :span="3" class="table-cell bold-text bg-dark"><a>End Date</a></el-col>
        <el-col :span="2" class="table-cell bold-text bg-dark"><a>Diff Day</a></el-col>
      </el-row>
  
      <el-row class="table" v-for="(item, index) in filteredData" :key="index">
        <el-col :span="3"></el-col>
        <el-col :span="2" class="table-cell"><a>{{ item.stt }}</a></el-col>
        <el-col :span="2" class="table-cell"><a>{{ item.buyNo }}</a></el-col>
        <el-col :span="3" class="table-cell"><a>{{ item.gsbh }}</a></el-col>
        <el-col :span="3" class="table-cell"><a>{{ item.ddbh }}</a></el-col>
        <el-col :span="3" class="table-cell"><a>{{ item.startDate }}</a></el-col>
        <el-col :span="3" class="table-cell"><a>{{ item.endDate }}</a></el-col>
        <el-col :span="2" class="table-cell"><a>{{ item.diffDay }}</a></el-col>
      </el-row>
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from "vue";
  import axios from "axios";
  
  const input = ref<string>("");
  const gsbhValue = ref("Tất cả");
  const procurementData = ref<any[]>([]);
  const filteredData = ref<any[]>([]); // Dữ liệu đã lọc
  const averageDiffDay = ref(0);
  
  const fetchProcurementData = async () => {
    try {
      const response = await axios.get("http://localhost:8080/api/procurement");
      procurementData.value = response.data;
      filteredData.value = procurementData.value; // Mặc định hiển thị toàn bộ dữ liệu
      calculateAverageDiffDay();
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };
  
  const calculateAverageDiffDay = () => {
    const totalDays = procurementData.value.reduce(
      (acc, item) => acc + item.diffDay,
      0
    );
    averageDiffDay.value = totalDays / procurementData.value.length;
  };
  
  // Hàm tìm kiếm với điều kiện: Nếu input rỗng, sẽ trả về tất cả dữ liệu
  const searchData = () => {
    filteredData.value = procurementData.value.filter((item) => {
      const matchesBuyNo = item.buyNo.toLowerCase().includes(input.value.toLowerCase());
      const matchesGsbh = gsbhValue.value === "Tất cả" || item.gsbh === gsbhValue.value;
      return matchesBuyNo && matchesGsbh;
    });
  };
  
  // Hàm reset search khi nhấn vào dấu x
  const resetSearch = () => {
    input.value = "";
    gsbhValue.value = "Tất cả"; // Reset GSBH về "Tất cả"
    filteredData.value = procurementData.value; // Trả về tất cả dữ liệu khi search bị xóa
  };
  
  onMounted(() => {
    fetchProcurementData();
  });
  </script>
  
  <style scoped>
  * {
    font-family: Arial, Helvetica, sans-serif;
  }
  
  .flex-start {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }
  
  .flex-end {
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }
  
  .input {
    margin-left: 10px;
    height: 32px;
    font-size: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    width: 100%;
  }
  
  .table {
    text-align: left;
    font-size: 18px;
    padding: 1px;
    background-color: #f5f5f5;
  }
  
  .table-cell a {
    display: block;
    padding: 8px;
    color: #333;
    text-decoration: none;
    text-align: center;
    background-color: #fff;
    border: 1px solid #ccc;
    border-radius: 4px;
    transition: background-color 0.3s ease, color 0.3s ease;
  }
  
  .table-cell a:hover {
    background-color: #f0f0f0;
    color: #007bff;
  }
  
  .bold-text {
    font-weight: bold;
  }
  
  .bg-dark {
    background-color: #d6d6d6;
  }
  
  /* Đảm bảo màu nền kéo dài từ đầu tới cuối trang */
  .table,
  .table-cell {
    background-color: #f5f5f5;
  }
  </style>
  