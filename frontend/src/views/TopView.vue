<template>
  <div id="top_page">
    <!-- header -->
    <div id="header">
      <div class="dummy"></div>
      <div id="title">
        <a id="main_title" href="/">{{ serviceName }}</a>
        <span id="sub_title">
          You can copy all formulas in an
          <a href="https://arxiv.org/" target="_blank">Arxiv</a> paper.</span
        >
      </div>
      <div id="sns_icons">
        <iframe
          src="https://ghbtns.com/github-btn.html?user=raahii&repo=arxiv-formula-extractor&type=star&count=true&size=large"
          frameborder="0"
          scrolling="0"
          width="170"
          height="30"
          title="GitHub"
        ></iframe>
      </div>
    </div>
    <div id="main">
      <div class="wrapper">
        <!--  search box -->
        <div id="search_box">
          <input v-model="arxivUrl" placeholder="https://arxiv.org/abs/..." />
          <button v-on:click="search" :disabled="isLoading">Go</button>
        </div>

        <!-- error message -->
        <p class="error" v-if="error">
          {{ error }}
        </p>

        <!-- rendering paper -->
        <div id="result">
          <PulseLoader class="loading_spinner" :loading="isLoading" />
          <Paper :obj="paper" v-if="!isLoading && paper" />
        </div>
      </div>
    </div>

    <!-- footer -->
    <div id="footer">
      {{ serviceName }} by
      <a :href="authorUrl" target="_blank">{{ authorName }}</a>
      <i class="fab fa-angellist"></i> | Contact:
      <a
        href="https://twitter.com/messages/compose?recipient_id=3304034184&ref_src=twsrc%5Etfw"
        class="twitter-dm-button"
        data-screen-name="@piyo56_net"
        data-show-count="false"
        target="_blank"
        >DM</a
      >
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Paper from "@/components/Paper.vue";
import PulseLoader from "vue-spinner/src/PulseLoader";

export default {
  name: "TopView",
  components: {
    Paper,
    PulseLoader,
  },
  data() {
    return {
      arxivUrl: "",
      paper: null,
      error: null,
      isLoading: false,

      // constants
      serviceName: "Arxiv Formula Extractor",
      authorName: "raahii",
      authorUrl: "https://raahii.github.io/about/",
      baseUrl: "https://arxiv.org/abs/",

      regex: new RegExp("https?://arxiv.org/abs/([0-9.]+)(v[0-9]+)?$"),
    };
  },
  mounted: function () {
    var arxivId = this.$route.query.arxiv_id;
    if (arxivId) {
      this.arxivUrl = this.baseUrl + arxivId;
      this.search();
    }
  },
  methods: {
    search: function () {
      this.error = null;
      if (this.checkUrl()) {
        this.findPaper();
      }
    },
    checkUrl: function () {
      let m = this.arxivUrl.match(this.regex);
      if (m === null) {
        if (this.arxivUrl.indexOf(this.baseUrl) == -1) {
          this.error = "The url must start 'https://arxiv.org/abs/'";
        } else {
          this.error = "Invalid arxiv url.";
        }
        return false;
      }

      return true;
    },
    setUrlParam: function () {
      this.$router.push({ query: { arxiv_id: this.arxivId } });
    },
    findPaper: function () {
      this.isLoading = true;
      this.paper = null;

      let self = this;
      axios({
        baseURL: this.apiUrl,
        method: "GET",
        url: "/papers/" + self.arxivId,
      })
        .then((response) => {
          self.paper = response.data.paper;
          self.setUrlParam();
        })
        .catch((error) => {
          let errorMsg = error.response.data.code + ": ";
          errorMsg += error.response.data.message;
          self.error = errorMsg;
        })
        .finally(() => {
          self.isLoading = false;
        });
    },
  },
  computed: {
    arxivId: function () {
      let m = this.arxivUrl.match(this.regex);
      return m[1];
    },
    apiUrl: function () {
      return process.env.VUE_APP_API_URL;
    },
  },
};
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
  background: #f5f5f5;
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

  @media screen and (min-width: 1000px) {
    width: 70%;
  }
  @media screen and (min-width: 1200px) {
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
    padding-left: 80px;
  }

  @media screen and (min-width: 700px) {
    height: $header-height-pc;
    .dummy {
      flex: 10%;
    }
    #title {
      flex: 50%;
    }
    #sns_icons {
      flex: 10%;
      justify-content: space-between;
      padding-left: 0px;
    }
  }
}
#main {
  padding: 20px 0 20px;
  background: white;

  height: auto;
  width: 100%;

  min-height: calc(100% - #{$header-height-sp} - #{$footer-height-sp});
  @media screen and (min-width: 700px) {
    min-height: calc(100% - #{$header-height-pc} - #{$footer-height-pc});
  }

  .wrapper {
    width: 90%;
    margin: 0 auto;

    @media screen and (min-width: 700px) {
      width: 80%;
    }
    @media screen and (min-width: 1000px) {
      width: 70%;
    }
    @media screen and (min-width: 1200px) {
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
        border: 1px solid #d1d7e3;
        border-radius: 4px;
        line-height: 26px;
        padding: 8px 36px 8px 14px;
        box-shadow: 0 4px 12px -2px rgba(#6b75a1, 0.16);
        color: #797c86;

        &::-webkit-input-placeholder {
          color: #c7c8cc;
        }
        &:-moz-placeholder {
          color: #c7c8cc;
        }
        &::-moz-placeholder {
          color: #c7c8cc;
        }
        &:-ms-input-placeholder {
          color: #c7c8cc;
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
  background: #f5f5f5;
  text-align: center;
  padding-top: calc(#{$footer-height-sp}/ 2.4);

  @media screen and (min-width: 700px) {
    height: $footer-height-pc;
    padding-top: calc(#{$footer-height-pc}/ 2.4);
  }
}
</style>
