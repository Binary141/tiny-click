<template>
  <div class="about">
    <h3 style="text-align: center">Add a URL</h3>
    <form>
        <p>
            URL: 
            <label for="scheme"></label>
            <select id="scheme" name="scheme">
                <option value="http">http://</option>
                <option value="https">https://</option>
            </select>
            <input style='margin:7px' type='text' id='newURL'>
        </p>
    </form>
    <button @click="buttonClick">Submit</button>
  </div>

</template>

<script>
    const serverURL = import.meta.env.VITE_SERVER_URL
    export default {
        methods: {
            async buttonClick(event) {
                var url = document.getElementById('newURL').value
                var scheme = document.getElementById('scheme').value

                var resp = await(fetch(`${serverURL}/insert?redirectURL=${scheme}://${url}`))
                console.log(resp);
                console.log("button was clicked")
            }
        }
    }
</script>
