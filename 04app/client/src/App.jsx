import FlightList from "./flights/FlightList"
import FlightCreate from "./flights/FlightCreate"
import FlightEdit from "./flights/FlightEdit"
import { BrowserRouter,Route,Routes } from 'react-router-dom'
import Counter from "./Counter"
import FullName from "./FullName"
function App() {//app is a functional component of UI
  return (
    <>
      <div>
        <BrowserRouter>
          <Routes>
            <Route path="" element={<FlightList/>}/>
            <Route path="/flights/list" element={<FlightList/>}/>
            <Route path="/flights/create" element={<FlightCreate/>}/>
            <Route path="/flights/edit/:id" element={<FlightEdit/>}/> //id path variable 
          </Routes>
        </BrowserRouter>
      </div>
    
      {/* <FullName/> */}
    </>
  )
}

export default App