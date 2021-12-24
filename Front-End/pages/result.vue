<template>
  <div class="container">


    <div v-if="type === 'phone'">
      <div v-if="history">
        <Phone :history="this.history" :search_id="this.search_id" :phone="this.keyword"/>
      </div>
      <div v-else>
        <Phone :phone="keyword"/>
      </div>

    </div>

    <div v-if="type === 'email'">
      <div v-if="history">
        <Email :history="this.history" :search_id="this.search_id" :email="this.keyword"/>
      </div>
      <div v-else>
        <Email :email="keyword"/>
      </div>

    </div>


    <div v-if="type === 'domain'">
      <div v-if="history">
        <Domain :history="this.history" :search_id="this.search_id" :domain="this.keyword"/>
      </div>
      <div v-else>
        <Domain :domain="keyword"/>
      </div>

    </div>

      <div v-if="type === 'name'">

        <div v-if="history">
          <Name :history="this.history" :search_id="this.search_id" :name="this.keyword"/>
        </div>
        <div v-else>
          <Name :name="keyword"/>
        </div>

        </div>




        <div v-if="type === 'ip'">

          <div v-if="history">
            <Ip :history="this.history" :search_id="this.search_id" :ip="this.keyword"/>
          </div>
          <div v-else>
            <Ip :ip="keyword"/>
          </div>

          </div>



    <div v-if="type === 'username'">

      <div v-if="history">
        <Username :history="this.history" :search_id="this.search_id" :username="this.keyword"/>
      </div>
      <div v-else>
        <Username :username="keyword"/>
      </div>

    </div>




  </div>
</template>

<script>

import Phone from "../components/Phone";
import Email from "../components/Email";
import Domain from "../components/Domain";
import Name from "../components/Name";
import Ip from '../components/Ip'
import Username from "../components/Username";
//import {toast} from "../helpers/toaster";


export default {

  // props : ["history" , "search_id" , "keyword" , "type"],
  components: {Username, Ip, Phone, Email, Domain , Name},


  created() {


    //alert(this.type)
    if (!(this.type) || !(this.keyword)) {

      //let that = this
      //toast(  that,"Insufficient data","error" , "top-right" , 3000)

      this.$router.push("/")
    }

  },
  data() {
    return {
      type: this.$route.params.type,
      keyword: this.$route.params.keyword,
      history: this.$route.params.history,
      search_id: this.$route.params.search_id
    }
  }
}

</script>

<style scoped>

</style>
