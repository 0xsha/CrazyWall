<template>
  <div>
    <a
      href="#"
      class=" mx-4 p-4 block bg-blue-dark hover:bg-blue-darker no-underline text-white border-b-2 border-gray-600 flex justify-between"
      @click.prevent="active = !active"
    >
      <strong class="text-gray-200">{{ title }}</strong>
      <span class="down-Arrow" v-show="!active">&#9660;</span>
      <span class="up-Arrow" v-show="active">&#9650;</span>
    </a>
    <div class="p-2" v-show="active">
      <slot />
    </div>
  </div>
</template>

<script>
export default {
  props: ["title", "objects"],
  data() {
    return {
      active: false
    };
  }
};
</script>

<style scoped></style>
