import { Button, TextArea } from "@patternfly/react-core";
import * as openpgp from "openpgp";
import { useState } from "react";

export function App() {
  const [generatedKey, setGeneratedKey] = useState("");

  return (
    <div>
      <Button
        variant="primary"
        onClick={async () => {
          const key = await openpgp.generateKey({
            userIds: [
              {
                name: "Felicitas Pojtinger",
                email: "felicitas@pojtinger.com",
              },
            ],
            passphrase: "123456",
          });

          setGeneratedKey(key.privateKeyArmored);
        }}
      >
        Generate key
      </Button>

      {generatedKey !== "" && (
        <TextArea
          value={generatedKey}
          aria-label="Generated PGP key"
          rows={generatedKey.split("\n").length}
        />
      )}
    </div>
  );
}
