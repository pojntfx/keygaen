import { CodeEditor, Language } from "@patternfly/react-code-editor";
import { Button } from "@patternfly/react-core";
import * as openpgp from "openpgp";
import React, { useState } from "react";

export function App() {
  const [generatedKey, setGeneratedKey] = useState("");

  return (
    <>
      <Button
        variant="primary"
        onClick={async () => {
          const key = await openpgp.generateKey({
            userIds: [
              {
                name: "Felix Pojtinger",
                email: "felix@pojtinger.com",
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
        <CodeEditor
          isReadOnly
          isMinimapVisible
          code={generatedKey}
          language={Language.plaintext}
          height="400px"
        />
      )}
    </>
  );
}
