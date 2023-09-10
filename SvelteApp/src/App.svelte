<!-- App.svelte -->
<script>
    import Home from "./Home.svelte";
    import Response from "./Response.svelte";

    let currentFeeling = null;
    let responseComponent;
    let timeout;

    function handleFeelingSelected(event) {
        currentFeeling = event.detail;
        clearTimeout(timeout);
        timeout = setTimeout(() => {
            resetFeeling();
        }, 20000);
    }

    function resetFeeling() {
        currentFeeling = null;
    }

    function handleResponseClick() {
        resetFeeling();
    }
</script>

<div>
    {#if currentFeeling === null}
        <Home on:feelingSelected={handleFeelingSelected} />
    {:else}
        <Response feeling={currentFeeling} bind:this={responseComponent} on:click={handleResponseClick} />
    {/if}
</div>
