<template>

  <ul v-for="url in urls" :key="url.redirectKey" style="text-align: center">
    <div :id="url.redirectKey">
      <a :href="serverURL + '/' + url.redirectKey" target="_blank">
        {{ url.redirectKey }}
      </a> 
      -> {{ url.redirectValue }}
      <button class="green" @click="deleteURL(url.redirectKey)">
        Delete
      </button> 
      <button v-if="editingKey !== url.redirectKey" class="green" @click="editURL(url.redirectKey, url.redirectValue)">
        Edit
      </button> 
      <button v-if="editingKey === url.redirectKey" class="green" @click="cancel">
        Cancel
      </button> 
    </div>
  <div v-if="editingKey === url.redirectKey" style='text-align: center'>
    <form @submit.prevent="submit">
      <p style="display: inline; text-align: center">
        <label for="scheme"></label>
        <select v-model="newScheme" name="scheme">
          <option value="http">http://</option>
          <option value="https">https://</option>
        </select>
        <input style="margin:0.3rem" type="text" v-model="newURL" />
      </p>
      <button style="display:inline" class="green">Submit</button>
    </form>
  </div>
  </ul>
</template>

<script>
import { ref, onMounted } from 'vue';

export default {
  setup() {
    const urls = ref([]);
    const serverURL = import.meta.env.VITE_SERVER_URL;
    const editingKey = ref(null);
    const newURL = ref('');
    const newScheme = ref('http');

    async function getURLs () {
      const urlsResp = await fetch(`${serverURL}/all`);
      urls.value = await urlsResp.json();
    }

    // Fetch data on component mount
    onMounted(async () => {
        getURLs();
    });

    async function submit () {
      const url = newURL.value;
      const scheme = newScheme.value;
      const redirectKey = editingKey.value;

      const reqURL = `${serverURL}/update?redirectKey=${redirectKey}&newRedirectURL=${scheme}://${url}`;
      const updateResp = await fetch(reqURL, {
          method: "PUT"
      });

      editingKey.value = null;
      getURLs();
    }

    function cancel () {
      editingKey.value = null;
    }

    function editURL (redirectKey, redirectValue) {
      const isHttps = redirectValue.includes('https://')
      const isHttp = redirectValue.includes('http://')
      
      if (isHttps) {
        newScheme.value = 'https'
        redirectValue = redirectValue.split('https://')[1]
      } else if (isHttp) {
        newScheme.value = 'http'
        redirectValue = redirectValue.split('http://')[1]
      }

      newURL.value = redirectValue
      editingKey.value = redirectKey
    }

    async function deleteURL (redirectKey) {
      const deleteResp = await fetch(`${serverURL}/${redirectKey}`, {
          method: "DELETE"
      });

      getURLs();
    }

    return {
      urls,
      serverURL,
      deleteURL,
      editURL,
      submit,
      cancel,
      editingKey,
      newURL,
      newScheme
    };
  }
};
</script>

