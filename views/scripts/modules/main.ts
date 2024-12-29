import './style.css'
import typescriptLogo from '../typescript.svg'
import { setupCounter } from './counter.ts'

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div>
  <p style="display: inline-block;">
      Powered by: 
    </p>
<img src="${typescriptLogo}" class="logo vanilla" style="display: inline-block;" alt="TypeScript logo" />
      <div class="card">
      <button id="counter" type="button"></button>
    </div>
  </div>
`

setupCounter(document.querySelector<HTMLButtonElement>('#counter')!)
