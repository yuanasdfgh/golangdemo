<template>
  <view class="container">
    <view class="h-12 w-12"> 首页 </view>
  </view>
  <div class="p-6 max-w-sm mx-auto bg-white rounded-xl shadow-md flex items-center space-x-4">
    <div class="flex-shrink-0">
      <img class="h-12 w-12" src="#" alt="ChitChat Logo" />
    </div>
    <div>
      <div class="text-xl font-medium text-red-700">ChitChat</div>
      <p class="text-gray-500">You have a new message!</p>
    </div>
  </div>
</template>
<script setup lang="ts"></script>

<style lang="scss" scoped>
.container {
  height: 185px;
  width: 375px;
  background: rgba(0, 0, 0, 0.1);
}
</style>
