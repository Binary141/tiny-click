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
      <button class="green" @click="editURL(url.redirectKey)">
        Edit
      </button> 
    </div>
  </ul>

  <div v-if="editingKey">
    <form @submit.prevent="submit">
      <p style="display: inline; text-align: center">
        <label for="scheme"></label>
        <select v-model="newScheme" name="scheme">
          <option value="http">http://</option>
          <option value="https">https://</option>
        </select>
        <input style="margin:0.3rem" type="text" v-model="newURL" />
      </p>
      <button type="submit" style="display:inline">Submit</button>
      <button @click="cancel" type="button" style="display:inline">Cancel</button>
    </form>
  </div>
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

    function editURL (redirectKey) {
      editingKey.value = redirectKey;
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

