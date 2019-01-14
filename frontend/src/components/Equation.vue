<template>
  <div v-on:mouseleave="onMouseLeave" class="equation card">
    <vue-mathjax 
      class="expression" 
      :formula="mathExp"
      v-clipboard:copy="expressionWithMacros"
      v-clipboard:success="onCopy"
      v-clipboard:error="onError"></vue-mathjax>
    <p class="copy_label">
    <span v-show="copy"><i class="far fa-clipboard"></i>copy</span>
    <span v-show="!copy"><i class="fas fa-check"></i>copied</span>
    </p>
  </div>
</template>

<script>
import { VueMathjax } from 'vue-mathjax'

export default {
  name: 'Paper',
  components: {
   'vue-mathjax': VueMathjax
  },
  data () {
    return {
      copy: true,
    }
  },
  props: ['eq', 'macros'],
  mounted: function () {
  },
  methods: {
    onCopy: function (e) {
      this.copy = false
    },
    onError: function (e) {
      console.error("copy error")
    },
    onMouseLeave: function (e) {
      this.copy = true
    }
  },
  computed: {
    mathExp: function () {
      let exp;
      exp  = String.raw`$$ \begin{align}`
      exp += this.eq.expression
      exp += String.raw`\end{align} $$`
      return exp
    },
    expressionWithMacros: function () {
      let macroString = ""
      for(let macro of this.macros) {
        if (this.eq.expression.indexOf(macro.command) >= 0) {
          macroString += macro.expression + "\n"
        }
      }

      return macroString + this.eq.expression
    },
    labelText: function () {
      if (this.copy) {
        return "copy"
      } else {
        return "copied!"
      }
    }
  }
}
</script>

<style scoped lang="scss">
.equation {
  width: 100%;
  margin-bottom: 20px;
  padding: 20px 0;
  position: relative;
}

.copy_label {
  display: flex;
  justify-content: center;
  align-items: center;

  visibility: hidden;
  opacity: 0;
  transition: .3s linear;

  position: absolute;
  top: 0;
  right: 0;

  padding: 0 5px;
  margin: 0;

  outline: none;
  border-top: none;
  border-right: none;
  border-left: solid 1px #D1D7E3;
  border-bottom: solid 1px #D1D7E3;
  line-height: 24px;
  box-shadow: 0 4px 12px -2px rgba(#6B75A1, .16);

  i {
    margin-right: 4px;
    color: #797C86;
  }

  .fa-clipboard {
    font-size: 15px;
  }

  .fa-check {
    font-size: 13px;
    color: #42b983;
  }

  span {
    font-size: 14px;
  }
}

.card {
  border: 1px solid;
  border-color: #cccccc;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  transition: all 0.3s cubic-bezier(.25,.8,.25,1);
}

.card:hover {
  cursor: pointer;
  box-shadow: 0 5px 5px rgba(0,0,0,0.25), 0 5px 5px rgba(0,0,0,0.22);
}

.card:hover .copy_label {
  visibility: visible;
  opacity: 1;
}
</style>
