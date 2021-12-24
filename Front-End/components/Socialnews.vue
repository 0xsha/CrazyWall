<template>
  <div class="text-gray-400">
    <div v-if="$fetchState.pending">
      <Preloader/>
    </div>
    <p v-else-if="$fetchState.error">Error while fetching live data please retry later</p>
    <div v-else>


      <div v-if="paging">
        <div
          class="fixed top-0 left-0 right-0 bottom-0 w-full h-screen z-50 overflow-hidden bg-gray-800 opacity-85 flex flex-col items-center justify-center">

          <img class="animate-pulse h-32 w-32 mb-4" src="~/assets/icons/coffee.svg" alt="coffee">
          <h2 class="text-center text-white text-xl font-semibold">Loading ...</h2>
<!--          <p class="w-1/3 text-center text-white">Hold on tight</p>-->


        </div>
      </div>


      <div class="flex justify-center items-center py-4" @click="searchSwitch">
        <span class="dark:text-draculaGreen px-2 font-bold " ref="switchText">{{ searchType }}</span>
        <div class="w-16 h-6 flex items-center bg-gray-300 rounded-full p-1 duration-300 ease-in-out"
             :class="{ 'bg-red-600': toggleActive}">
          <div class="bg-white w-4 h-4 rounded-full shadow-md transform duration-300 ease-in-out"
               :class="{ 'translate-x-10': toggleActive,}"></div>
        </div>
      </div>


      <div v-if="this.socialSlice && this.toggleActive">


        <!--          <div v-for='postArray in this.fullPosts'  >-->


        <div v-for="(post , index) in this.socialSlice">


          <div class="rounded-lg bg-draculaBackground py-4 mb-2 mt-2 px-4">


            <a :href="post['url']"  >
              <img :src="require(`~/assets/icons/${post['network']}.svg`)"  class="h-12 w-12 " alt="">
            </a>

            <div class="flex items-center ">


            <p >
            {{ post["text"] }}
            </p>
            </div>
<!--            {{ post["url"] }}-->
<!--            {{ post["image"] }}-->
<!--            {{ post["user"]["url"] }}-->
            <div class="flex justify-between pt-2">
              <p class="text-draculaGreen">{{ post["posted"] }}</p>

              <a :href='post["user"]["url"]'>

             <span class="text-draculaPurple"> {{ post["user"]["name"] }}</span>
              </a>

            </div>

          </div>
        </div>


      </div>

    </div>


    <div v-if='this.newsSlice && !this.toggleActive'>


      <div class="bg-gray-900 w-full h-8"></div>


      <div class="bg-draculaBackground rounded-lg">

        <div v-for='article in this.newsSlice'>


          <a :href='article["url"]' target="_blank">
            <div class="text-draculaForeground text-xl text-center bg-draculaPurple py-2 rounded-lg">
              {{ article["title"] }}
            </div>
          </a>

          <div class="grid grid-cols-2 gap-4 py-4 ">


            <div class="px-4">
              <a :href='article["url"]' target="_blank">


                <img v-if="article['urlToImage']" :src='article["urlToImage"]'
                     class="rounded-lg object-cover h-72 w-full" alt="img">
                <img v-else src="~assets/images/404.webp" class="rounded-lg object-cover h-72 w-full" alt="img">

              </a>
            </div>

            <div class="px-4">

              {{ article["description"] }}
            </div>


            <div class="flex flex-col md:flex-row px-2  text-draculaGreen md:text-xl md:justify-evenly">
              <span> {{ article["source"]["name"] }}</span>
              <span>{{ article["publishedAt"] }} </span>
            </div>

          </div>


          <div class="bg-gray-900 w-full h-8"></div>

        </div>

      </div>


    </div>

  </div>


</template>

<script>
import Preloader from "./Preloader";

export default {
  components: {Preloader},
  props: ["keyword"],
  name: "Socialnews",


  data() {

    return {
      newsSlice: {},
      fullPosts: [],
      socialSlice: [],
      paging: false,
      offset: 0,
      sOffset: 0,
      searchType: 'News',
      toggleActive: false,

    }
  },

  async fetch() {


    await this.endScroll()


    let details = await this.$axios.post("/api/v1/live/search", {
      keyword: this.keyword
    })


    if (details.data["social"].length >= 1) {


      for (let i = 0; i < details.data["social"].length; i++) {

        this.fullPosts.push(details.data["social"][i]["FullData"][0]["posts"])

      }

      this.fullPosts = this.fullPosts.flat()
      if (this.fullPosts.length > 20) {
        this.socialSlice = this.fullPosts.slice(0, 20)
        this.sOffset = 20
      } else {

        this.socialSlice = this.fullPosts
      }

    }


    this.news = details.data["news"]["articles"]

    if (details.data["news"]["articles"].length > 20) {
      this.newsSlice = details.data["news"]["articles"].slice(0, 20)
      this.offset = 20
    } else {
      this.newsSlice = details.data["news"]["articles"]
    }


  },

  methods: {



    searchSwitch() {
      this.toggleActive = !this.toggleActive


      if (this.toggleActive) {
        // this.$refs.searchInput.placeholder = "keyword (across social networks and news)"
        this.$refs.switchText.textContent = "Social"
      } else {
        //  this.$refs.searchInput.placeholder = "domain, email, phone, name, username, IP"
        this.$refs.switchText.textContent = "News"
      }

      this.searchType = this.$refs.switchText.textContent


    },

    async endScroll() {

      // check if user scrolled to end of the page
      window.onscroll = async () => {
        let bottomOfWindow = Math.max(window.pageYOffset, document.documentElement.scrollTop, document.body.scrollTop) + window.innerHeight === document.documentElement.offsetHeight


        if (bottomOfWindow) {


          // hack ?
          if (this.$route.name === "live") {


            if (this.searchType === "News") {


              if (this.offset >= this.news.length) {

                //this.offset = this.news.length
                this.newsSlice = this.news
                this.paging = false

                return
              }


              await clearTimeout(this.timeout);
              this.paging = true

              this.timeout = await setTimeout(() => {

                if ((this.offset + 20) > this.news.length) {

                  this.offset = this.news.length - this.offset
                  this.newsSlice = this.news.slice(0, this.offset)
                  this.paging = false

                }

                if ((this.offset + 20) <= this.news.length) {

                  this.offset += 20

                  this.newsSlice = this.news.slice(0, this.offset)


                }
                this.paging = false


              }, 3000)

            }


            if (this.searchType === "Social") {

              if (this.sOffset >= this.fullPosts.length) {

                //this.offset = this.news.length
                this.socialSlice = this.fullPosts
                this.paging = false

                return
              }


              await clearTimeout(this.timeout);

              this.paging = true

              this.timeout = await setTimeout(() => {


                if ((this.sOffset + 20) > this.fullPosts.length) {

                  this.sOffset = this.fullPosts.length - this.sOffset
                  this.socialSlice = this.fullPosts.slice(0, this.offset)
                  this.paging = false

                }

                if ((this.sOffset + 20) <= this.fullPosts.length) {

                  this.sOffset += 20

                  this.socialSlice = this.fullPosts.slice(0, this.sOffset)


                }
                this.paging = false


              }, 3000)

            }


          }


        }


      }

    }


  },

  destroyed() {
    window.removeEventListener('scroll', () => {
    })
  },
}


</script>

<style scoped>

</style>
