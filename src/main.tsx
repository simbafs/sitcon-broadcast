import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { App } from './App'
import './main.css'

createRoot(document.getElementById('root')!).render(
	<StrictMode>
		<div className="flex flex-col items-center justify-center h-screen w-screen">
			<App />
		</div>
	</StrictMode>,
)
