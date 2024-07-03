<script setup>
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
