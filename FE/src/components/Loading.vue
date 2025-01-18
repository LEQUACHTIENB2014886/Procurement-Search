<template>
  <transition name="fade" @after-leave="onFadeOutComplete">
    <div v-if="!isCompleted" class="loading-overlay">
      <div class="circle-container">
        <svg width="120" height="120" viewBox="0 0 120 120">
          <circle
            class="background-circle"
            cx="60"
            cy="60"
            r="50"
            fill="none"
            stroke="#e6e6e6"
            stroke-width="10"
          />
          <circle
            class="progress-circle"
            cx="60"
            cy="60"
            r="50"
            fill="none"
            :stroke="circleColor"
            stroke-width="10"
            stroke-dasharray="314.16"
            :stroke-dashoffset="314.16 - (314.16 * progress) / 100"
          />
        </svg>

        <div class="progress-text">{{ progress }}</div>
      </div>
      <p>{{ text }}</p>
    </div>
  </transition>

</template>

<script setup>
import { ref, computed, onMounted, defineProps } from "vue";

const props = defineProps({
  text: {
    type: String,
    default: "Loading...",
  },
});

const isCompleted = ref(false);
const progress = ref(0);

const gradientColors = [
  { color: "#e74c3c", stop: 0 },
  { color: "#f39c12", stop: 50 },
  { color: "#2ecc71", stop: 100 },
];

const circleColor = computed(() => {
  let currentColor = "#e74c3c";
  for (let i = 0; i < gradientColors.length - 1; i++) {
    const start = gradientColors[i];
    const end = gradientColors[i + 1];
    if (progress.value >= start.stop && progress.value <= end.stop) {
      const factor = (progress.value - start.stop) / (end.stop - start.stop);
      currentColor = interpolateColor(start.color, end.color, factor);
      break;
    }
  }
  return currentColor;
});

function hexToRgb(hex) {
  const bigint = parseInt(hex.replace("#", ""), 16);
  const r = (bigint >> 16) & 255;
  const g = (bigint >> 8) & 255;
  const b = bigint & 255;
  return { r, g, b };
}

function interpolateColor(startColor, endColor, factor) {
  const start = hexToRgb(startColor);
  const end = hexToRgb(endColor);
  const r = Math.round(start.r + factor * (end.r - start.r));
  const g = Math.round(start.g + factor * (end.g - start.g));
  const b = Math.round(start.b + factor * (end.b - start.b));
  return `rgb(${r}, ${g}, ${b})`;
}

onMounted(() => {
  const interval = setInterval(() => {
    if (progress.value < 100) {
      progress.value += 1;
    } else {
      clearInterval(interval);
      setTimeout(() => {
        isCompleted.value = true; 
      }, 1000); 
    }
  }, 10);
});


const onFadeOutComplete = () => {
  // Callback function for when the fade out animation is complete
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Oswald:wght@400;500;600&display=swap");

.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.8);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  z-index: 9999;
  color: white;
  font-size: 18px;
  font-family: "Oswald", sans-serif; /* Sử dụng font Oswald */
}

.circle-container {
  position: relative;
  width: 120px;
  height: 120px;
}

.background-circle {
  stroke: #e6e6e6;
}

.progress-circle {
  transform: rotate(-90deg);
  transform-origin: center;
  stroke-linecap: round;
  transition: stroke-dashoffset 0.3s;
}

.progress-text {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: 36px;
  font-weight: 600;
  color: white;
  letter-spacing: 3px; 
}

.completed-status {
  text-align: center;
  font-size: 28px; 
  color: white;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
