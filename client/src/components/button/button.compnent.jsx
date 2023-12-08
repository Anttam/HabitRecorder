import "./button.styles.css";
import { APIURL } from "../../api";
import goodHabitButton from '../../assets/good habit button.png'
import badHabitButton from '../../assets/bad habit button.png'


const Button = ({ type }) => {
  const handleButtonClick = async () => {
    try {
      const res = await fetch(APIURL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ habit: type, uid: "test"} )
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
      <button className=" button button-up" onClick={handleButtonClick}>
        <img src={goodHabitButton} className="button-img"/>
      </button>

    );
  if (type === "bad") {
    return (
    <button className="button button-down"onClick={handleButtonClick}>
        <img src={badHabitButton} className="button-img"/>
       </button>)
  }
};

export default Button;
