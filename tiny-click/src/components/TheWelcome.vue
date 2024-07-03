<script setup>
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import ToolingIcon from './icons/IconTooling.vue'
import EcosystemIcon from './icons/IconEcosystem.vue'
import CommunityIcon from './icons/IconCommunity.vue'
import SupportIcon from './icons/IconSupport.vue'
</script>

<template>
    <ul v-for="url in urls" style="text-align: center">
        <a :href="'http://localhost:5000/' + url.redirectKey" target=_blank>{{ url.redirectKey }}</a> -> {{ url.redirectValue }}
    </ul>

</template>

<script>
    import { ref } from 'vue'
    var urlsResp = await(fetch('http://localhost:5001/all'))

    const urls = ref(await(urlsResp.json()))
</script>
