<template>
  <div class="top_page">
    <!-- header -->
    <div id="header">
      <div id="title">
        <a href="/">{{ service_name }}</a>
      </div>
      <div id="subtitle">
        provides latex format equations from <a href="https://arxiv.org/">Arxiv</a>.
      </div>
    </div>
  
    <div id="main">
      <!--  search box -->
      <div id="search_box">
        <input v-model="arxiv_url" placeholder="https://arxiv.org/abs/...">
        <button v-on:click="search" v-bind:disabled="isLoading">Go</button>
      </div>

      <!-- error message -->
      <p id="error" v-if="error">
        {{ error }}
      </p>

      <!-- rendering paper -->
      <div id="result">
        <pulse-loader class='loading_spinner' :loading="isLoading"></pulse-loader>
        <paper v-bind:obj="paper" v-if="paper && !isLoading"></paper>
      </div>
    </div>

    <!-- footer -->
    <div id="footer">
      {{ service_name }} by <a :href="author_url" target="_blank">{{ author_name }}</a>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Paper from './Paper.vue'
import PulseLoader from "vue-spinner/src/PulseLoader";

export default {
  name: 'TopPage',
  components: {
    "paper": Paper,
    "pulse-loader": PulseLoader,
  },
  data () {
    return {
      service_name: 'Arxiv Equations',
      service_url: 'http://localhost:8000',
      author_name: 'raahii',
      author_url: 'https://raahii.github.io/about/',
      url_prefix: "https://arxiv.org/abs/",
      arxiv_url: '',
      paper: null,
      error: null,
      isLoading: false,
    }
  },
  mounted: function() {
    var arxiv_id = this.$route.query.arxiv_id
    if (arxiv_id) {
      this.arxiv_url = this.url_prefix + arxiv_id
      this.find_paper()
    }
  },
  methods: {
    search: function(e) {
      this.error = null
      if (this.checkUrl(e)) {
        this.find_paper()
      }
    },
    checkUrl: function () {
      if (this.arxiv_url.indexOf(this.url_prefix) != 0) {
        this.error = "The url must start 'https://arxiv.org/abs/'"
        return false
      }
      return true
    },
    setUrlParam: function() {
      this.$router.push({query: {arxiv_id: this.arxiv_id}})
    },
    find_paper: function () {
      this.isLoading = true
      this.paper = null
      
      let self = this
      axios({
          method : 'GET',
          url    : '/papers/' + self.arxiv_id,
        })
        .then(response => {
          self.paper = response.data.paper
          self.isLoading = false
          self.setUrlParam()
        })
        .catch(error => {
          let errorMsg = error.response.data.code + ": "
          errorMsg += error.response.data.message
          self.error = errorMsg
          self.isLoading = false
        })
    },
  },
  computed: {
    arxiv_id: function () {
      let idStr = this.arxiv_url.slice(this.url_prefix.length)
      let pos = idStr.indexOf("v")
      if (pos != -1) {
        idStr = idStr.slice(0, pos)
      }
      return idStr
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
p {
  font-weight: normal;
  font-size: 24px;
  text-align: left;
}
a {
  color: #42b983;
}
#header {
  width: 100%;
  height: auto;
  background: #F5F5F5;
  padding: 20px 0;
  text-align: center;

  #title {
    font-size: 32px;
    a {
      color: #2c3e50;
      text-decoration: none;
    }
  }
}
#main {
  min-height: 100vh;
  height: 100%;
  margin: 40px auto 0;

  width: 90%;
  @media screen and (min-width:700px) { 
    width: 80%;
  }
  @media screen and (min-width:1000px) { 
    width: 60%;
  }

  #search_box {
    width: 100%;
    margin: 20px auto 20px;
    display: flex;
    flex-direction: row;
    align-content: center;

    input {
      flex: 0 1 80%;
      display: block;
      box-sizing: border-box;
      color: #2c3e50;
      height: 40px;
      padding: 5px 5px;
      font-size: 18px;
      border-color: gray;
    }
    button {
      flex: 0 1 20%;
      display: block;
      box-sizing: border-box;
      height: 40px;
      margin-left: 5px;
      line-height: 40px;

      background-color: #42b983;
      border: none;
      color: white;
      text-align: center;
      text-decoration: none;
      display: inline-block;
      font-size: 18px;
    }
  }

  #result {
    width: 100%;
    margin: 20px 0;
    p {
      font-size: 14px;
    }
    .loading_spinner {
      width: 100%;
      margin: 0 auto;
      text-align: center;
    }
  }
}

#footer {
  box-sizing: border-box;
  height: 100px;
  background: #F5F5F5;
  padding-top: 45px;
  text-align: center;
  margin-top: 30px;
}
</style>
