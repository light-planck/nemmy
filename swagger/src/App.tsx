import SwaggerUI from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

const App = () => {
  return (
    <div className="App">
      <SwaggerUI url="openapi.yaml" />
    </div>
  );
};

export default App;
