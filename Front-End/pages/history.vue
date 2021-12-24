<template>
  <div class="container">
    <div class="  py-4 text-gray-50 bg-gray-900 mt-4 py-4 rounded-lg ">

      <div class="py-8  ">
        <div>
          <h2 class="text-2xl font-semibold leading-tight dark:text-draculaForeground">History</h2>
        </div>
        <div class="my-2 flex sm:flex-row flex-col">
          <div class="flex flex-row mb-1 sm:mb-0">
            <div class="relative ">
              <select v-model="selected" @change="perPage"
                      class="appearance-none  h-full  rounded-l dark:bg-gray-800 dark:text-gray-300 dark:border-draculaComment border block appearance-none w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
                <option disabled value="">...</option>
                <option>5</option>
                <option>10</option>
                <option>20</option>
              </select>
              <div
                class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700 dark:text-gray-300">
                <svg class="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
                  <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"/>
                </svg>
              </div>
            </div>
          </div>
          <div class="block relative">
                    <span class="h-full absolute inset-y-0 left-0 flex items-center pl-2">
                        <svg viewBox="0 0 24 24" class="h-4 w-4 fill-current text-gray-400">
                            <path
                              d="M10 4a6 6 0 100 12 6 6 0 000-12zm-8 6a8 8 0 1114.32 4.906l5.387 5.387a1 1 0 01-1.414 1.414l-5.387-5.387A8 8 0 012 10z">
                            </path>
                        </svg>
                    </span>
            <input v-model="keyword" placeholder="Search" @keyup="search"
                   class="appearance-none rounded-r rounded-l sm:rounded-l-none dark:bg-gray-800 border border-gray-400 dark:border-draculaComment border-b block pl-8 pr-6 py-2 w-full bg-white text-sm placeholder-gray-400 text-gray-700 dark:text-gray-300 focus:bg-white focus:placeholder-gray-600 focus:text-gray-700 focus:outline-none"/>
          </div>
        </div>
        <div class="-mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto   ">
          <div class="inline-block min-w-full shadow rounded-lg overflow-hidden">
            <table v-if="!loading" class="min-w-full leading-normal   ">
              <thead>
              <tr>
                <th
                  class="px-5 py-3 border-b-2 border-gray-200 dark:border-draculaComment bg-gray-100 dark:bg-gray-700 dark:text-draculaGreen text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                  Keyword
                </th>
                <th
                  class="px-5 py-3 border-b-2 border-gray-200 dark:border-draculaComment bg-gray-100 dark:bg-gray-700 dark:text-draculaGreen text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                  Date
                </th>
                <th
                  class="px-5 py-3 border-b-2 border-gray-200 dark:border-draculaComment bg-gray-100 dark:bg-gray-700 dark:text-draculaGreen text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                  Type
                </th>
                <th
                  class="px-5 py-3 border-b-2 border-gray-200 dark:border-draculaComment bg-gray-100 dark:bg-gray-700 dark:text-draculaGreen text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                  View
                </th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="history in historyArray">
                <td
                  class="px-5 py-5 border-b border-gray-200  dark:border-draculaComment bg-white text-sm dark:bg-gray-700 dark:text-draculaGreen ">
                  <div class="flex items-center">
                    <div class="ml-3">
                      <p class="text-gray-900 dark:bg-gray-700 dark:text-draculaForeground whitespace-no-wrap"
                         ref="historyKeyword">
                        {{ history["keyword"] }}
                      </p>
                    </div>
                  </div>
                </td>

                <td
                  class="px-5 py-5 border-b dark:bg-gray-700  border-gray-200 dark:border-draculaComment bg-white text-sm">
                  <p class="text-gray-900 dark:text-draculaForeground  whitespace-no-wrap">
                    {{ formatDate(history["created_at"]) }}</p>
                </td>
                <td
                  class="px-5 py-5 border-b border-gray-200 dark:bg-gray-700 dark:border-draculaComment bg-white text-sm">
                  <p class="text-gray-900 dark:text-draculaForeground whitespace-no-wrap">
                    {{ history["s_type"] }}
                  </p>
                </td>
                <td
                  class="px-5 py-5 border-b border-gray-200 dark:bg-gray-700  dark:border-draculaComment bg-white text-sm">

                  <a @click="getResults(history['id'] ,  history['s_type'] , history['keyword'])">
                    <img class="h-10 w-10 rounded-full " src="~/assets/icons/eye.svg" alt="">

                  </a>


                </td>
              </tr>

              </tbody>
            </table>

            <div v-if="loading" v-for="_ in 5" class="flex justify-center py-4">
              <div class=" bg-gray-400 w-full h-10  rounded-lg animate-pulse"></div>
            </div>

            <div v-if="historyArray"
                 class="px-5 py-5 bg-white dark:text-draculaGreen dark:bg-gray-700 border-t dark:border-draculaComment flex flex-col xs:flex-row items-center xs:justify-between          ">

              <!--                        pagination-->
              <span class="text-xs xs:text-sm text-gray-900 dark:text-draculaForeground">
                            Showing {{
                  this.offset + 1
                }} to {{ (this.limit + this.offset) >= this.total ? this.total : (this.limit + this.offset) }} of {{
                  this.total
                }} Entries
                        </span>

              <div class="inline-flex mt-2 xs:mt-0">
                <button @click="pervPage"
                        :class="{'dark:hover:bg-black':this.hasPrev , 'cursor-not-allowed' : !this.hasPrev}"
                        :disabled="!this.hasPrev"
                        class="text-sm bg-gray-300 hover:bg-gray-400 disabled:cursor-not-allowed disabled:opacity-50  dark:bg-draculaBackground dark:text-draculaPurple   text-gray-800 font-semibold py-2 px-4 rounded-l">
                  Prev
                </button>
                <button @click="nextPage"
                        :class="{'dark:hover:bg-black':this.hasNext , 'cursor-not-allowed' : !this.hasNext}"
                        :disabled="!this.hasNext"
                        class=" disabled:opacity-50 text-sm bg-gray-300 page-button hover:bg-gray-400 text-gray-800 dark:bg-draculaBackground dark:text-draculaPurple font-semibold py-2 px-4 rounded-r">
                  Next
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>


    </div>
  </div>
</template>

<script>
export default {

  data() {
    return {

      historyArray: [],
      hasPrev: false,
      hasNext: false,
      limit: 0,
      offset: 0,
      total: 0,
      search_id: 0,
      loading: false,
      selected: "",
      keyword: ""


    }
  },

  async created() {

    this.$axios.onError(error => {
      this.loading = false
    })

    this.$axios.onRequest(req => {

      this.loading = true
    })

    this.$axios.onResponse(response => {


      clearTimeout(this.timeout)

      this.timeout = setTimeout(() => {

        this.total = response.data["total"]
        this.historyArray = response.data["history"]


        // pagination buttons
        this.hasNext = (this.offset + this.limit) <= this.total
        this.hasPrev = ((Math.abs(this.limit - this.offset) >= 0) && (this.offset !== 0))


        this.loading = false

      }, 300)


    })


  },

  // skeleton loading on search
  async mounted() {


    // reset pagination
    this.limit = 5
    this.offset = 0

    await this.$axios.post("/api/v1/history/mine", {limit: this.limit, offset: this.offset})


  },

  methods: {

    getResults(searchId, hType, hKeyword) {

      this.$router.push({
        name: "result",
        params: {search_id: searchId, type: hType, history: "true", keyword: hKeyword}
      })
    },

    // simple date formatter
    formatDate(d) {


      let date = new Date(d)

      const monthNames = ["January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"
      ]
      return monthNames[date.getMonth()] + " " + date.getDate() + "," + date.getFullYear()


    },

    async nextPage() {

      if (this.offset + this.limit >= this.total) {
        return
      } else {

        this.offset += this.limit
      }

      if (this.keyword !== "") {
        await this.$axios.post("/api/v1/history/mine", {
          limit: this.limit,
          offset: this.offset,
          action: "search",
          keyword: this.keyword
        })
        return
      }

      await this.$axios.post("/api/v1/history/mine", {limit: this.limit, offset: this.offset})


    },

    async pervPage() {


      if (this.offset <= 0) {
        return
      } else {

        this.offset -= this.limit

      }


      if (this.keyword !== "") {
        await this.$axios.post("/api/v1/history/mine", {
          limit: this.limit,
          offset: this.offset,
          action: "search",
          keyword: this.keyword
        })
        return
      }

      await this.$axios.post("/api/v1/history/mine", {limit: this.limit, offset: this.offset})


    },

    async perPage() {

      // set new limit
      this.limit = parseInt(this.selected)


      if (this.keyword !== "") {
        await this.$axios.post("/api/v1/history/mine", {
          limit: parseInt(this.selected),
          action: "search",
          keyword: this.keyword
        })
        return
      }
      await this.$axios.post("/api/v1/history/mine", {limit: parseInt(this.selected)})

    },
    async search() {

      // clear timeout
      await clearTimeout(this.timeout);
      //let that = this;
      // inline debounce
      this.timeout = await setTimeout(async () => {

          this.offset = 0
          let limit = 5
          if (this.selected !== "") {

            limit = parseInt(this.selected)

          }

          await this.$axios.post("/api/v1/history/mine", {
            "action": "search",
            "keyword": this.keyword, limit: limit
          })
        }
        , 500)


    }
  }
}
</script>

<style scoped>
</style>
