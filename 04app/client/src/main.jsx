import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.jsx'

createRoot(document.getElementById('root')).render( // react will capture the snapshot of doc
  <StrictMode>
    <App />
  </StrictMode>,
)
