<template>
  <div class="top_page">
    <!-- header -->
    <div id="header">
      <div id="dummy"></div>
      <div id="title">
        <a id="main_title" href="/">{{ serviceName }}</a>
        <span id="sub_title"> provides latex format equations from <a href="https://arxiv.org/" target="_blank">Arxiv</a>.</span>
      </div>
      <div id="sns_icons">
        <a :href="githubUrl" target="_blank"><img id="github" src="/static/github.png"></a>
      </div>
    </div>
  
    <div id="main">
      <!--  search box -->
      <div id="search_box">
        <input v-model="arxivUrl" placeholder="https://arxiv.org/abs/...">
        <button v-on:click="search" v-bind:disabled="isLoading">Go</button>
      </div>

      <!-- error message -->
      <p class="error" v-if="error">
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
      {{ serviceName }} by <a :href="authorUrl" target="_blank">{{ authorName }}</a>
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
      serviceName: 'Arxiv Equations',
      serviceUrl: 'http://localhost:8000',
      authorName: 'raahii',
      authorUrl: 'https://raahii.github.io/about/',
      baseUrl: "https://arxiv.org/abs/",
      githubUrl: "https://github.com/raahii/arxiv-equations",

      baseUrl: "https://arxiv.org/abs/",
      regex: new RegExp("https?://arxiv.org/abs/([0-9.]+)(v[0-9]+)?$"),

      arxivUrl: '',
      paper: null,
      error: null,
      isLoading: false,
    }
  },
  mounted: function() {
    var arxivId = this.$route.query.arxiv_id
    if (arxivId) {
      this.arxivUrl = this.baseUrl + arxivId
      this.search()
    }
  },
  methods: {
    search: function(e) {
      this.error = null
      if (this.checkUrl()) {
        this.findPaper()
      }
    },
    checkUrl: function () {
      let m = this.arxivUrl.match(this.regex)
      if (m === null) {
        if (this.arxivUrl.indexOf(this.baseUrl) == -1) {
          this.error = "The url must start 'https://arxiv.org/abs/'"
        }  else {
          this.error = "Invalid arxiv url."
        }
        return false
      }
      
      return true
    },
    setUrlParam: function() {
      this.$router.push({query: {arxiv_id: this.arxivId}})
    },
    findPaper: function () {
      this.isLoading = true
      this.paper = null
      
      let self = this
      axios({
          method : 'GET',
          url    : '/papers/' + self.arxivId,
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
    arxivId: function () {
      let m = this.arxivUrl.match(this.regex)
      return m[1]
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
  display: flex;
  
  #dummy {
    flex: 33%;
  }
  #title {
    flex: 34%;
    display: inline-flex;
    flex-direction: column;
    margin: 10px 0;

    #main_title {
      display: inline-box;
      margin: 0 auto;
      font-size: 32px;
      color: #2c3e50;
      text-decoration: none;
    }
    #sub_title {
      display: inline-box;
      margin: 0 auto;
    }
  }
  #sns_icons {
    flex: 33%;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;

    #github {
      height: 32px;
      width: 32px;
      border-radius: 50%;
      box-shadow: 0 6px 12px -2px rgba(107, 117, 161, 0.16);
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

      width: 100%;
      height: 40px;

      padding: 5px 5px;
      font-size: 18px;
      outline: none;
      border: 1px solid #D1D7E3;
      border-radius: 4px;
      line-height: 26px;
      padding: 8px 36px 8px 14px;
      box-shadow: 0 4px 12px -2px rgba(#6B75A1, .16);
      color: #797C86;

      &::-webkit-input-placeholder {
        color: #C7C8CC;
      }
      &:-moz-placeholder {
        color: #C7C8CC;
      }
      &::-moz-placeholder {
        color: #C7C8CC;
      }
      &:-ms-input-placeholder {
        color: #C7C8CC;
      }
    }
    button {
      flex: 0 1 20%;
      display: block;
      box-sizing: border-box;
      height: 40px;
      margin-left: 5px;
      line-height: 40px;

      border-radius: 4px;
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
