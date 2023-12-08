import './about.styles.css'

const About = () => {
  return (
    <div className="about-container">
      <h2>What is the Habit Tracker?</h2>
      <p>
        The habit tracker is an idea I came up with after reading "Atmoic
        Habits" by James Clear. In the book, there is a section talking about
        the difference between having a good or bad day is the sum of the habits
        we perform throughout it. The idea is that if you follow multiple tiny
        habits (or "votes" as Clear describes it) you will gradually mold
        yourself into the person you envision your self as.
      </p>
      <h2>How does it work?</h2>
      <p>
        This app allows you to hit the button that corrisponds with the habit
        you performed (green = good red = bad). The app keeps track of your
        habits throughout the day and will give you a heads up at 9pm EST on
        whether you had a good or bad habit day.
      </p>
    </div>
  );
};

export default About;
