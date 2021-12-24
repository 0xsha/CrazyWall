<template>
  <div>
    <div v-if="$fetchState.pending">
      <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none"
           viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>

    </div>
    <p v-else-if="$fetchState.error">Error</p>
    <div v-else>
      <img :src="flagURL" class="w-12 h-6  rounded-l rounded-r">
    </div>
  </div>

</template>

<script>
export default {
  props: ['code'],
  data() {
    return {
      flagURL: ''
    }
  },

  async fetch() {

    // we use fetch here because we globally set axios to backend endpoint ;)

    // await fetch("https://restcountries.eu/rest/v2/alpha/" + this.code)
    //   .then(res => res.json())
    //   .then(data => {
    //     this.flagURL = data.flag
    //   })
    //   .catch(error => {
    //     console.log(error)
    //   })

    let flag = await this.$axios.post("/api/v1/auxiliary/flag", {code: this.code})
    this.flagURL = flag.data["flag"]
  }

}
</script>

<style scoped>

</style>
