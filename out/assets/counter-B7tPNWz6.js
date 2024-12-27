function c(n){let t=0;const e=o=>{t=o,n.innerHTML=`count is ${t}`};n.addEventListener("click",()=>e(t+1)),e(0)}export{c as setupCounter};
