import './app.styles.css'
import Button from "./components/button/button.compnent"
import Header from './components/header/header.component'


function App() {
  return (<>
    <div className='header-container'>
    <Header/>
   <div className='buttons-container'>
    <Button type='good'/>
    <Button type ='bad'/>
   </div>
    </div>
  </>
   
  )
}

export default App
