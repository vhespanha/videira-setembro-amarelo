<!-- Response.svelte -->
<script>
    export let feeling;
    import { createEventDispatcher } from "svelte";

    function getBackgroundColor(feeling) {
      const backgroundColors = {
        Felicidade: "#fbb83d",
        Tristeza: "#264b93",
        Medo: "#4c3492",
        Raiva: "#f02e35",
        Ansiedade: "#3ab066",
        Não_sei: "#87408f",
      };
      return backgroundColors[feeling] || "#FFFFFF"; // Default to white if feeling is not found
    }

    const backgroundColor = getBackgroundColor(feeling);

    function getRandomResponse(feeling) {
      // Defina suas respostas para cada emoção aqui
      const responses = {
        Felicidade: [
          "Celebre esse sentimento fazendo algo que você ama!",
          "Sorria! Compartilhe com alguém esse sentimento bom.",
          "Você merece muitos momentos felizes!",
          "Que tal dançar para celebrar esse sentimento?",
          "Você está radiante! Celebre!",
        ],
        Tristeza: [
          "Tudo é temporário. Respire fundo, esse momento difícil vai passar",
          "Seja gentil consigo mesmo, você está dando o seu melhor.",
          "Tire um tempo para fazer algo que você ama, se acolha, você merece.",
          "Parabéns pela sua coragem em continuar mesmo com os desafios da vida",
          "Sinto muito pelo que você está passando, eu vejo você, seu sentimento importa.",
        ],
        Medo: [
          "Tudo bem sentir medo, tudo bem recuar. você não precisa ter coragem o tempo todo.",
          "Lembre-se, você é capaz de fazer tudo que se dedicar.",
          "O medo faz parte da jornada, e você pode escolher continuar, recuar, ou mudar o caminho.",
          "Você não precisa ter todas as respostas. Respire fundo e vá um passo de cada vez.",
          "Lembre-se: você não precisa se preocupar com coisas que não consegue controlar.",
        ],
        Raiva: [
          "Tudo faz parte da nossa jornada. O que você está aprendendo com esse momento?",
          "Tudo é temporário. Respire fundo, esse momento difícil vai passar.",
          "Que tal parar, respirar e acolher o que você está sentindo?",
          "Pegue papel e caneta e escreva seus pensamentos, prometo que tudo ficará mais leve.",
          "Respire. Se afastar do que causa esse sentimento pode ajudar a melhorar.",
        ],
        Ansiedade: [
          "Lembre-se: você não precisa se preocupar com coisas que não consegue controlar.",
          "Pegue papel e caneta e escreva seus pensamentos, prometo que tudo ficará mais leve.",
          "Faça exercício físico ou pratique uma atividade que traga sua mente para o presente.",
          "Respire contando até 4, expire contando até 4. Repita durante 1 minuto, Vai melhorar",
          "Você não precisa ter todas as respostas. Respire fundo e vá um passo de cada vez",
        ],
        Não_sei: [
          "A psicoterapia é uma ótima ferramente para entender, acolher e cuidar dos seus sentimentos.",
          "Orgulhe-se da sua jornada! Cada passo, pequeno ou grande, acrescenta algo em sua vida",
          "Que tal pagar uma folha de papel e descarregar tudo que está passando na sua cabeça? Este é um ótimo exercício para esclarecer os pensamentos e as emoçoes",
          "Respire fundo durante 1 minuto. Qualquer que seja o sentimento vai melhorar",
          "Que tal iniciar o hábito de meditação? Observe e deixe suas emoções fluirem",
        ],
      };
      const responseList = responses[feeling];
      if (responseList) {
        const randomIndex = Math.floor(Math.random() * responseList.length);
        return responseList[randomIndex];
      }
      return "No response found for this feeling.";
    }

    const selectedResponse = getRandomResponse(feeling);

    const dispatch = createEventDispatcher();

    function handleClick() {
      dispatch("click");
    }
</script>

<div class="content-wrapper">
    <div class="overlay" style="background-color: {backgroundColor};" />

    <h1 class="response">{selectedResponse}</h1>
    <button type="button" class="reload" on:click={handleClick} />
</div>
<div class="logos">
    <img src="logo1.png" alt="Logo 1" />
</div>


<style>
    .reload {
        background-image: none;
        background-color: rgba(0, 0, 0, 0);
        height: 100%;
        width: 100% !important;
        position: absolute;
        top: 0%;
        left: 0%;
        box-shadow: none;

    }
    .content-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100vh; /* Isso garante que o conteúdo ocupe toda a altura da janela */
    }
    .response {
        line-height: 1;
        text-align: center;
        animation: zoom-in-animation 2s ease-in-out;
        color: #0f0f0f;
        font-size: 40px;
    }
    .overlay {
        position: absolute; /* Para posicionar a overlay dentro do elemento pai */
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: white; /* Defina a cor e opacidade desejadas aqui */
        z-index: -1; /* Certifique-se de que a overlay esteja acima da imagem de fundo */
        pointer-events: none; /* Para permitir cliques nos elementos subjacentes */
        background-image: url("overlay_bg.png");
        background-size: 100%;
    }

    @keyframes zoom-in-animation {
        0% {
            transform: scale(0.8); /* Começa médio */
            opacity: 0; /* Começa invisível, se desejar */
        }
        100% {
            transform: scale(1); /* Fica grande */
            opacity: 1; /* Fica totalmente visível */
        }
    }
    .logos {
        display: flex;
        position: absolute;
        top: 80%;
    }

    .logos img {
        width: 150pt; /* Defina o tamanho desejado para as logos */
        height: auto; /* Mantém a proporção da imagem */
        border-radius: 20px;
        box-shadow: 0px 0px 30px 0px rgba(0, 0, 0, 0.5);
        position: fixed;
        top: 93%;
        left: 50%;
        transform: translate(-50%, -50%);
    }

</style>
