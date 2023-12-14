import SwaggerUI from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

const App = () => {
  // const origin = import.meta.env.VITE_APP_ORIGIN;

  return (
    <div className="App">
      {/* <SwaggerUI url={`${origin}/doc/openapi.yaml`} /> */}
      <SwaggerUI url={`.vercel/output/static/openapi.yaml`} />
    </div>
  );
};

export default App;
