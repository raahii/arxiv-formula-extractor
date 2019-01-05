<template>
  <div class="main">
    <!--  title, description -->
    <div id="upper">
      <div id="title">
        {{ service_name }}
      </div>
      <div id="subtitle">
        provides latex format equations from <a href="https://arxiv.org/">Arxiv</a>
      </div>
    </div>

    <!--  search box -->
    <div id="search_box">
      <input v-model="arxiv_url" placeholder="enter paper url">
      <button v-on:click="find_paper">Go</button>
    </div>

    <!-- rendering paper -->
    <div id="result" v-if="searched">
      <paper v-bind:obj="paper"></paper>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Paper from './Paper.vue'

export default {
  name: 'TopPage',
  components: {
    Paper
  },
  data () {
    return {
      service_name: 'Arxiv Equations',
      arxiv_url: '',
      searched: false,
      paper: {},
    }
  },
  methods: {
    find_paper: function () {
      axios
        .get("http://localhost:1323/papers?url="+this.arxiv_url)
        .then(response => {
          if (response.status != 200)  {
            console.log(response.error)
            this.errored = true
            return
          }

          this.searched = true
          this.paper = response.data.paper
        })
        .catch(error => {
          console.log(error)
          this.errored = true
        })
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
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#upper {
  width: 100%;
  height: auto;
  background: #FAFAFA;
  padding: 20px 0;
  text-align: center;

  #title {
    font-size: 32px;
  }
}
#search_box {
  width: 80%;
  margin: 20px auto;
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
#result {
  p {
    font-size: 14px;
  }
}
</style>
