import "./button.styles.css";
import { APIURL } from "../../api";

const Button = ({ type }) => {
  const handleButtonClick = async () => {
    
    let habitValue = type === "good" ?  1 :  -1
    try {
      const res = await fetch(APIURL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ number: habitValue} )
      });

      console.log(res.body)

      if (!res.ok) {
        throw new Error(`error, status: ${res.status}`);
      }
    } catch (err) {
      console.log(err);
    }
  };

  if (type === "good")
    return (
      <button className="button-up" onClick={handleButtonClick}>
        
        &#9650;
      </button>
    );
  if (type === "bad") {
    return <button className="button-down"onClick={handleButtonClick}> &#9660;</button>;
  }
};

export default Button;
