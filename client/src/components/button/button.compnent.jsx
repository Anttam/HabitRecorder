import './button.styles.css'

const Button = ({type}) =>{
  if(type === 'good')
return(
  <button className="button-up"> &#9650;</button>
)
if (type === 'bad'){
  return(
    <button className="button-down"> &#9660;</button>
  )
}
}

export default Button