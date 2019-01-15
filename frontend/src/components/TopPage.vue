<template>
  <div id="top_page">
    <!-- header -->
    <div id="header">
      <div class="dummy"></div>
      <div id="title">
        <a id="main_title" href="/">{{ serviceName }}</a> <span id="sub_title"> provides latex format equations from <a href="https://arxiv.org/" target="_blank">Arxiv</a>.</span>
      </div>
      <div id="sns_icons">
        <!-- <a :href="githubUrl" target="_blank"><img id="github" src="/static/github.png"></a> -->
        <a class="github-button" href="https://github.com/raahii/arxiv-equations" data-size="large" data-show-count="true" aria-label="Star raahii/arxiv-equations on GitHub">Star</a>
      </div>
    </div>
    <div id="main">
      <div class="wrapper">
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
          <paper :obj="paper" v-if="!isLoading && paper"></paper>
        </div>
      </div>
    </div>

    <!-- footer -->
    <div id="footer">
       {{ serviceName }} by <a :href="authorUrl" target="_blank">{{ authorName }}</a> <i class="fab fa-angellist"></i> |
       Contact: <a href="https://twitter.com/messages/compose?recipient_id=3304034184&ref_src=twsrc%5Etfw" class="twitter-dm-button" data-screen-name="@piyo56_net" data-show-count="false" target="_blank">DM</a>
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
          self.setUrlParam()
        })
        .catch(error => {
          let errorMsg = error.response.data.code + ": "
          errorMsg += error.response.data.message
          self.error = errorMsg
        })
        .then(()=> {
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
#top_page {
  height: 100%;
  background: #F5F5F5;
}

.dummy {
  visibility: hidden;
}

$header-height-sp: 150px;
$footer-height-sp: 70px;
$header-height-pc: 120px;
$footer-height-pc: 70px;

#header {
  box-sizing: border-box;
  padding: 20px 0;
  margin: 0 auto;
  display: flex;
  flex-wrap: wrap;

  height: $header-height-sp;
  width: 90%;
  margin: 0 auto;

  @media screen and (min-width:700px) { 
    width: 80%;
  }
  @media screen and (min-width:1000px) { 
    width: 60%;
  }

  .dummy {
    flex: 0%;
  }

  #title {
    flex: 100%;
    display: inline-flex;
    flex-direction: column;
    margin: 10px 0;
    text-align: center;

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
    flex: 100%;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
    #github {
      height: 24px;
      width: 24px;
      border-radius: 50%;
      box-shadow: 0 6px 12px -2px rgba(107, 117, 161, 0.16);
    }
  }
  
  @media screen and (min-width:700px) { 
    height: $header-height-pc;
    .dummy {
      flex: 25%;
    }
    #title {
      flex: 50%;
    }
    #sns_icons {
      flex: 25%;
      #github {
        height: 32px;
        width: 32px;
      }
    }
  }
}
#main {
  padding: 20px 0 20px;
  background: white;

  height: auto;
  width: 100%;

  min-height: calc(100% - #{$header-height-sp} - #{$footer-height-sp});
  @media screen and (min-width:700px) { 
    min-height: calc(100% - #{$header-height-pc} - #{$footer-height-pc});
  }
  
  .wrapper {
    width: 90%;
    margin: 0 auto;

    @media screen and (min-width:700px) { 
      width: 80%;
    }
    @media screen and (min-width:1000px) { 
      padding-top: 20px;
      width: 60%;
    }

    #search_box {
      width: 100%;
      margin: 0 auto;
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
}

#footer {
  box-sizing: border-box;
  height: $footer-height-sp;
  line-height: 0px;
  background: #F5F5F5;
  text-align: center;
  padding-top: calc(#{$footer-height-sp/2.4);

  @media screen and (min-width:700px) { 
    height: $footer-height-pc;
    padding-top: calc(#{$footer-height-pc}/2.4);
  }
}
</style>
