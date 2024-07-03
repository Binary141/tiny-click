<template>
  <div class="about">
    <h3 style="text-align: center">Add a URL</h3>
  </div>

<button @click="buttonClick">Submit</button>
</template>

<script>
    export default {
  methods: {
    buttonClick(event) {
      console.log("button was clicked");
    }
  }
}
</script>
