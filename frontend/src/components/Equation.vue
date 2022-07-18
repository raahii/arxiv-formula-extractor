<template>
  <div v-on:mouseleave="onMouseLeave" class="equation card">
    <vue-mathjax
      class="expression"
      :formula="renderExpression"
      v-clipboard:copy="copyExpression"
      v-clipboard:success="onCopy"
      v-clipboard:error="onError"
    ></vue-mathjax>
    <p class="copy_label">
      <span v-show="copy"><i class="far fa-clipboard"></i>copy</span>
      <span v-show="!copy"><i class="fas fa-check"></i>copied</span>
    </p>
  </div>
</template>

<script>
import { VueMathjax } from "vue-mathjax";

export default {
  name: "Paper",
  components: {
    "vue-mathjax": VueMathjax,
  },
  data() {
    return {
      copy: true,
    };
  },
  props: ["eq", "macros"],
  methods: {
    onCopy: function () {
      this.copy = false;
    },
    onError: function () {
      console.error("copy error");
    },
    onMouseLeave: function () {
      this.copy = true;
    },
  },
  computed: {
    renderExpression: function () {
      let exp = "";
      exp += String.raw`$$ \begin{align}` + "\n";
      exp += this.eq.expression + "\n";
      exp += String.raw`\end{align} $$`;
      return exp;
    },
    copyExpression: function () {
      let exp = "";
      if (this.macros.length > 0) {
        exp += this.macros.map((m) => m.expression).join("\n") + "\n\n";
      }
      exp += String.raw`\begin{eqnarray}` + "\n";
      exp += this.eq.expression + "\n";
      exp += String.raw`\end{eqnarray}`;

      return exp;
    },
    labelText: function () {
      if (this.copy) {
        return "copy";
      } else {
        return "copied!";
      }
    },
  },
};
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
  transition: 0.3s linear;

  position: absolute;
  top: 0;
  right: 0;

  padding: 0 5px;
  margin: 0;

  outline: none;
  border-top: none;
  border-right: none;
  border-left: solid 1px #d1d7e3;
  border-bottom: solid 1px #d1d7e3;
  line-height: 24px;
  box-shadow: 0 4px 12px -2px rgba(#6b75a1, 0.16);

  i {
    margin-right: 4px;
    color: #797c86;
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
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
}

.card:hover {
  cursor: pointer;
  box-shadow: 0 5px 5px rgba(0, 0, 0, 0.25), 0 5px 5px rgba(0, 0, 0, 0.22);
}

.card:hover .copy_label {
  visibility: visible;
  opacity: 1;
}
</style>
