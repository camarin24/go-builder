<script lang="ts">
    let value: string;
    let dbType: string;
    let ok: boolean;

    async function ping() {
        const res = await fetch("http://127.0.0.1:8051/api/admin/ping", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                connStr: value,
                connType: dbType,
            }),
        });

        if (res.ok) {
            const response = await res.json();
            ok = response.ping;
        }
    }
</script>

<main>
    <div class="card">
        <h3>Create your first project</h3>
        <select name="" id="" bind:value={dbType} placeholder="Choose your database">
            <option value="pg">PostgreSQL</option>
            <option value="mysql">Mysql</option>
            <option value="sqlite">Sqlite</option>
        </select>
        <input type="text" bind:value name="" placeholder="String connection" id="" />
        <button on:click={ping}>Test connection</button>
        <br />
        {#if ok != undefined}
            {#if ok}
                <span class="success">Connection successfully</span>
            {:else}
                <span class="error">Connection error</span>
            {/if}
        {/if}
    </div>
</main>

<style>
    .success {
        color: #9ade7b;
        border: 1px solid #9ade7b;
    }
    .error {
        color: #ff8f8f;
        border: 1px solid #ff8f8f;
    }
</style>
