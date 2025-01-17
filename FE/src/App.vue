<template>
  <div id="app">
    <loading v-if="isLoading" text="Loading..." />
    <div v-else>
      <Table />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import Loading from "./components/Loading.vue";
import Table from "./components/Table.vue";

const isLoading = ref(true); //

onMounted(() => {
  setTimeout(() => {
    isLoading.value = false;
  }, 1500);
});
</script>

<style>
#app {
  text-align: center;
  margin-top: 50px;
}
</style>