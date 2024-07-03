<template>
    <ul v-for="url in urls" style="text-align: center">
        <a :href="serverURL + '/' + url.redirectKey" target=_blank>
            {{ url.redirectKey }}
        </a> 
        -> {{ url.redirectValue }}
        <button class='green' @click='deleteURL(url.redirectKey)'>
            Delete
        </button> 

    </ul>

</template>

<script>
import { ref, onMounted } from 'vue';

export default {
  setup() {
    const urls = ref([]);
    const serverURL = import.meta.env.VITE_SERVER_URL;

    async function getURLs () {
      const urlsResp = await fetch(`${serverURL}/all`);
      urls.value = await urlsResp.json();
    }

    // Fetch data on component mount
    onMounted(async () => {
        getURLs()
    });

    async function deleteURL (redirectKey) {
      console.log(redirectKey)
      const deleteResp = await fetch(`${serverURL}/${redirectKey}`, 
          {
              method: "DELETE"
          }
      );

      getURLs()
    }

    return {
      urls,
      serverURL,
      deleteURL
    };
  }
};
</script>
