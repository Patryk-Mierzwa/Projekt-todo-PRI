<template>
  <div id="app">
    <b-container>
     <h3>Aplikacja - Lista zadań</h3>
     <h3></h3>
     <h3>Wprowadź zadanie</h3>
      <div>
        <TodoApp />
      </div>
    </b-container>
  </div>
</template>

<script>
import TodoApp from "./components/TodoApp.vue";

export default {
  components: {
    TodoApp,
  },
};
</script>

<style> 
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
}
</style>
