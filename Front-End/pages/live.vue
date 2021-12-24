<template>
  <div class="container">

    <div v-if="keyword">
    <Socialnews :keyword="keyword"></Socialnews>
    </div>

  </div>



</template>

<script>
import Socialnews from "../components/Socialnews";


export default {
  components : {Socialnews},
  name: "live",

  created() {

    if (!this.keyword) {

       this.$router.push("/")

    }
  },

  data(){
    return {

      keyword: this.$route.params.keyword

    }
  }

}
</script>

<style scoped>

</style>
