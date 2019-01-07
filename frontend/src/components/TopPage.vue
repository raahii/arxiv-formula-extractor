<template>
  <div class="top_page">
    <!-- header -->
    <div id="header">
      <div id="title">
        <router-link to="/">{{ service_name }}</router-link>
      </div>
      <div id="subtitle">
        provides latex format equations from <a href="https://arxiv.org/">Arxiv</a>.
      </div>
    </div>
  
    <div id="main">
      <!--  search box -->
      <div id="search_box">
        <input v-model="arxiv_url" placeholder="https://arxiv.org/abs/...">
        <button v-on:click="find_paper" v-bind:disabled="isLoading">Go</button>
      </div>

      <!-- error message -->
      <p id="errors" v-if="errors.length">
        <ul>
          <li class='error' v-for="error in errors">{{ error }}</li>
        </ul>
      </p>

      <!-- rendering paper -->
      <div id="result">
        <pulse-loader class='loading_spinner' :loading="isLoading"></pulse-loader>
        <paper v-bind:obj="paper" v-if="searched && !isLoading"></paper>
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
      paper: {},
      errors: [],
      searched: false,
      isLoading: false,
    }
  },
  mounted: function() {
    var arxiv_id = this.$route.query.arxiv_id
    console.log(arxiv_id)
    if (arxiv_id) {
      this.arxiv_url = this.url_prefix + arxiv_id
      this.find_paper()
    }
  },
  methods: {
    search: function(e) {
      this.checkUrl(e)
      this.find_paper()
    },
    checkUrl: function (e) {
      this.errors = [];

      if (this.arxiv_url.indexOf(this.url_prefix) == -1) {
        this.errors.push("The url must start 'https://arxiv.org/abs/'");
      }
      e.preventDefault();
    },
    setParam: function() {
      let parts = this.arxiv_url.split("/")
      this.$router.push({query: {arxiv_id: parts[parts.length-1]}})
    },
    find_paper: function () {
      this.isLoading = true
      this.searched = false
      this.setParam()

      axios
        .get("http://localhost:1323/papers?url="+this.arxiv_url)
        .then(response => {
          if (response.status != 200)  {
            console.log(response.error)
            this.errored = true
            return
          }

          this.paper = response.data.paper
        })
        .catch(error => {
          console.log(error)
        })
        .finally(() => {
          this.searched = true
          this.isLoading = false
        })
    },
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
  width: 90%;
  min-height: 100vh;
  height: 100%;

  @media screen and (min-width:700px) { 
    width: 80%;
  }
  @media screen and (min-width:1000px) { 
    width: 60%;
  }
  margin: 0 auto;

  #search_box {
    width: 100%;
    margin: 20px auto 40px;
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
  #errors {
    width: 100%;
    font-size: 1.2em;
    margin-bottom: 3em;

    ul {
      list-style-type: disc;
    }
    .error {
     color:red;
   }
  }
  
  #result {
    width: 100%;
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
