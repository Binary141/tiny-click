<template>
    <ul v-for="url in urls" style="text-align: center">
        <a :href="'http://localhost:5000/' + url.redirectKey" target=_blank>{{ url.redirectKey }}</a> -> {{ url.redirectValue }}
    </ul>

</template>

<script>
import { ref, onMounted } from 'vue';

export default {
  setup() {
    const urls = ref([]);
    const serverURL = import.meta.env.VITE_SERVER_URL;

    // Fetch data on component mount
    onMounted(async () => {
        const urlsResp = await fetch(`${serverURL}/all`);
      urls.value = await urlsResp.json();
    });

    return {
      urls,
      serverURL
    };
  }
};
</script>
