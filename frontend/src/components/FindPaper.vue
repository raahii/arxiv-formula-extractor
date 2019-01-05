<template>
  <div class="hello">
    <input v-model="query" placeholder="edit me">
    <button v-on:click="search">Go</button>
    <div v-show="searched">
      <p>search result of "{{ query }}"</p>
      <p>{{ result.n_results }} hits!</p> 
      <ul>
        <paper v-for="(value, key) in result.papers" :paper="value" :key="key"></paper>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Paper from './Paper.vue'

export default {
  name: 'FindPaper',
  components: {
    Paper
  },
  data () {
    return {
      query: '',
      result: {},
      searched: false,
    }
  },
  methods: {
    search: function (event) {
      axios
        .get("http://localhost:1323/papers?q="+this.query)
        .then(response => {
          if (response.status != 200)  {
            console.log(response.error)
            this.errored = true
            return
          }
          
          this.searched = true
          this.result = response.data
        })
        .catch(error => {
          console.log(error)
          this.errored = true
        })
      }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  width: 600px;
  margin: auto;
  padding: 0;
  list-style-type: circle;
  text-align: left;
}
a {
  color: #42b983;
}
</style>
