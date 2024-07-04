<template>
  <div class="about">
    <h3 style="text-align: center; margin-bottom: 1rem">Add a URL</h3>
    <form style='text-align: center'>
        <p style='display: inline'>
            <label for="scheme"></label>
            <select id="scheme" name="scheme">
                <option value="http">http://</option>
                <option value="https">https://</option>
            </select>
            <input style='margin:0.3rem' type='text' id='newURL'>
        </p>
        <button @click="buttonClick" style='display:inline'>Submit</button>
    </form>
  </div>

</template>

<script>
    const adminServerURL = import.meta.env.VITE_ADMIN_SERVER_URL
    export default {
        methods: {
            async buttonClick(event) {
                event.preventDefault()
                var url = document.getElementById('newURL').value
                var scheme = document.getElementById('scheme').value

                var resp = await(fetch(`${adminServerURL}/insert?redirectURL=${scheme}://${url}`))
                console.log(resp);
                console.log("button was clicked")
            }
        }
    }
</script>
