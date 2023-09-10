<script>
    import { createEventDispatcher } from "svelte";

    const dispatcher = createEventDispatcher();

    function handleFeelingClick(feeling) {
      fetch("http://localhost:8081/save-feeling", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: `feeling=${feeling}`,
      })
        .then((response) => {
          if (response.ok) {
            console.log("Sentimento salvo com sucesso!");
          } else {
            console.error("Erro ao salvar sentimento");
          }
        })
        .catch((error) => {
          console.error("Erro ao conectar-se ao servidor", error);
        });
      dispatcher("feelingSelected", feeling);
    }
    let emotions = [
      { name: "Felicidade", gif: "felicidade.gif", class: "felicidade" },
      { name: "Tristeza", gif: "tristeza.gif", class: "tristeza" },
      { name: "Medo", gif: "medo.gif", class: "medo" },
      { name: "Raiva", gif: "raiva.gif", class: "raiva" },
      { name: "Ansiedade", gif: "ansiedade.gif", class: "ansiedade" },
      { name: "Não_sei", gif: "não_sei.gif", class: "não_sei" },
    ];
</script>

<div class="home-container">
    <h1>O que você está sentindo?</h1>
    <h2>Selecione e receba uma mensagem</h2>
    <div class="feeling-list">
        {#each emotions as emotion}
          <div class="button-container">
              <button
                  class={emotion.class}
                  on:click={() => handleFeelingClick(emotion.name)}
                  >
                  {emotion.name === "Não_sei" ? "Não sei" : emotion.name}
              </button>
              <div class="gif-container">
                  <img src={emotion.gif} alt={"GIF de " + emotion.name} />
              </div>
          </div>
        {/each}
      </div>
</div>
<div class="logos">
    <img src="logo1.png" alt="Logo 1" />
</div>

<style>
    .logos {
      display: flex;
      position: absolute;
      top: 80%;
    }

    .logos img {
      width: 150pt;
      height: auto;
      border-radius: 20px;
      box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.5);
      position: fixed;
      top: 93%;
      left: 50%;
      transform: translate(-50%, -50%);
    }

</style>
