<template>
  <div id="app">
    <loading v-if="isLoading" text="Loading..." />
    <Table v-else />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import Loading from "./components/Loading.vue";
import Table from "./components/Table.vue";

const isLoading = ref(true);

onMounted(() => {
  setTimeout(() => isLoading.value = false, 1300);
});
</script>

<style>
#app {
  text-align: center;
  margin-top: 50px;
}
*{
  font-family: "Oswald", sans-serif; 
}
</style>
