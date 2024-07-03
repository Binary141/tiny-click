<script setup>
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import ToolingIcon from './icons/IconTooling.vue'
import EcosystemIcon from './icons/IconEcosystem.vue'
import CommunityIcon from './icons/IconCommunity.vue'
import SupportIcon from './icons/IconSupport.vue'
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Documentation</template>

    Vue’s
    <a href="https://vuejs.org/" target="_blank" rel="noopener">official documentation</a>
    provides you with all information you need to get started.
  </WelcomeItem>
  <button @click="test">This is a test</button>

    <ul v-for="url in urls" style="text-align: center">
        {{ url.redirectKey }} -> {{ url.redirectValue }}
    </ul>

</template>

<script>
    import { ref } from 'vue'
    var urlsResp = await(fetch('http://localhost:5001/all'))
    // .then(function(response) {
    //     response.json().then(function (data) {
    //         testing = data
    //         console.log(data);
    //     })
    // })

    const urls = ref(await(urlsResp.json()))

    function test(event) {
        console.log("Hello there button!");
    }
</script>
