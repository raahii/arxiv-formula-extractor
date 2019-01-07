<template>
  <div class="top_page">
    <!--  title, description -->
    <div id="header">
      <div id="title">
        {{ service_name }}
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
        <paper v-bind:obj="paper" v-if="!isLoading"></paper>
      </div>
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
      arxiv_url: '',
      searched: false,
      paper: {},
      errors: [],
      isLoading: false,
    }
  },
  methods: {
    checkUrl: function (e) {
      this.errors = [];
      let prefix = "https://arxiv.org/abs/"

      if (this.arxiv_url.indexOf(prefix) == -1) {
        this.errors.push("The url must start 'https://arxiv.org/abs/'");
      }
      e.preventDefault();
    },
    find_paper: function (e) {
      this.checkUrl(e)

      this.isLoading = true
      this.searched = false
      this.paper = null

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
  background: #FAFAFA;
  padding: 20px 0;
  text-align: center;

  #title {
    font-size: 32px;
  }
}
#main {
  width: 90%;
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
      height: 30px;
      font-size: 18px;
      border-color: gray;
    }
    button {
      flex: 0 1 20%;
      display: block;
      box-sizing: border-box;
      height: 30px;
      margin-left: 2px;
      background: black;
      color: white;
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
</style>
