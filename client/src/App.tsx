import logo from "./assets/logo.svg";
import { Button } from "./components/ui/button";

function App() {
  return (
    <main className="">
      <img src={logo} className="w-80" />
      <Button>Click Here</Button>
    </main>
  );
}

export default App;
