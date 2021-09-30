import { CodeEditor, Language } from "@patternfly/react-code-editor";
import { Button as PFButton } from "@patternfly/react-core";
import * as openpgp from "openpgp";
import React, { useState } from "react";

export const Button = () => {
  const [generatedKey, setGeneratedKey] = useState("");

  return (
    <>
      <PFButton
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
      </PFButton>

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
};
