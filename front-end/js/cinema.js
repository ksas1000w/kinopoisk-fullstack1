const tr_img = document.getElementById("trailer-img")
const tr_video = document.getElementById("trailer-video")
const bigDescription = document.querySelector(".m_big_description")
const tr_logo = document.getElementById("trailer-logo")
const trailer_container = document.querySelector(".m_trailer_container")
const DELAY = 7000;
const h_btns = document.querySelectorAll(".h_btn")
const rightArrow = document.querySelector(".rightArrow")
const leftArrow = document.querySelector(".leftArrow")
// const videoSrc = ["Attack on Titan Season 1 Trailer [LV-nazLVmgo].f248.webm", "dandadan.mp4"]
// const imgSrc = ["//avatars.mds.yandex.net/get-ott/212840/2a00000186a8b7a1951185c6175ed0f07fd0/2016x1134", "//avatars.mds.yandex.net/get-ott/1652588/2a000001981bf0c61fd778cd9bc97146001c/2016x1134"]
// const logoSrc = ["//avatars.mds.yandex.net/get-ott/1531675/2a0000018e7eebec99e598e322d5f2ebc8cb/960x540", "//avatars.mds.yandex.net/get-ott/2419418/2a00000199baba1d2b377a01006f01546779/960x540"]
// const text = ["Люди сражаются с титанами, которые мечтают их съесть. Финал самого эпичного аниме современности", "Внучка медиума и юный уфолог внезапно встречают призраков и пришельцев. Хитовое аниме — безумное и смешное"]
// const films.length = 2

let films = []
let counter = 0

async function loadFilms() {
    const response = await fetch("/api/films")
    films = await response.json()
    console.log(films) 
}

loadFilms()

function isImg(newVideoSrc = null) {
    tr_video.pause();
    tr_video.classList.add('m_tc_opacityOn');
    tr_logo.classList.add('m_tc_opacityOn');
    tr_img.classList.remove('m_tc_opacityOn');
    bigDescription.classList.remove('m_tc_opacityOn');
    setTimeout(() => {
        tr_video.src = "";
        tr_video.load();
        if (newVideoSrc) {
            tr_video.src = newVideoSrc;
        }
    }, 1500);
}

function isVideo() {
    tr_video.play();
    tr_video.classList.remove('m_tc_opacityOn');
    tr_logo.classList.remove('m_tc_opacityOn');
    tr_img.classList.add('m_tc_opacityOn');
    bigDescription.classList.add('m_tc_opacityOn');
}

rightArrow.addEventListener("click", () => {
    observer.disconnect();
    counter++
    if (counter == films.length) counter = 0
    tr_img.src = films[counter].card.path;
    tr_logo.src = films[counter].logo.path;
    isImg(films[counter].trailer.path)
    bigDescription.innerHTML = `<img src="${films[counter].logo.path}"><p>${films[counter].description}</p>`
    observer.observe(trailer_container)
})

leftArrow.addEventListener("click", () => {
    observer.disconnect();
    counter--
    if (counter == -1) counter = films.length - 1
    tr_img.src = films[counter].card.path;
    tr_logo.src = films[counter].logo.path;
    isImg(films[counter].trailer.path)
    bigDescription.innerHTML = `<img src="${films[counter].logo.path}"><p>${films[counter].description}</p>`
    observer.observe(trailer_container)
})

const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        const card = entry.target;
        if (entry.isIntersecting) {
            card._timer = setTimeout(() => { isVideo() }, DELAY);
        } else {
            clearTimeout(card._timer);
            isImg()
        }
    });
}, { threshold: 0.5 });

observer.observe(trailer_container)

h_btns.forEach(btn => {
    btn.addEventListener("click", () => {
        for (const temp of h_btns)
            if (temp.classList.contains("h_select_btn")) {
                temp.classList.remove("h_select_btn")
                break
            }
        btn.classList.add("h_select_btn")
    })
})

// Горизонтальный скролл со стрелками
document.querySelectorAll('.am-teka-wrapper').forEach(wrapper => {
    const cont = wrapper.querySelector('.am-teka-cont');
    const btnLeft = wrapper.querySelector('.am-arrow-left');
    const btnRight = wrapper.querySelector('.am-arrow-right');
    const scrollAmount = 600;

    btnLeft.addEventListener('click', () => {
        cont.scrollBy({ left: -scrollAmount, behavior: 'smooth' });
    });
    btnRight.addEventListener('click', () => {
        cont.scrollBy({ left: scrollAmount, behavior: 'smooth' });
    });
});