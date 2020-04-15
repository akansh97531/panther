import React from 'react';
import { Box, Card, Grid, Text } from 'pouncejs';
import Panel from 'Components/Panel';
import SQLEditor from './SQLEditor';
import Browser from './Browser';
import Results from './Results';
import { useSQLShellContext, withSQLShellContext } from './SQLShellContext';

const SQLShellPage: React.FC = () => {
  const {
    state: { globalErrorMessage },
  } = useSQLShellContext();

  return (
    <Box>
      {globalErrorMessage && (
        <Text
          size="large"
          color="red300"
          is="p"
          textAlign="center"
          backgroundColor="red50"
          p={6}
          mb={4}
        >
          {globalErrorMessage}
        </Text>
      )}
      <Grid gridGap={4} gridTemplateColumns="1fr 3fr" mb={4}>
        <Browser />
        <Card p={9}>
          <SQLEditor />
        </Card>
      </Grid>
      <Box mb={4} minHeight={400}>
        <Panel title="Results" size="large">
          <Results />
        </Panel>
      </Box>
    </Box>
  );
};

export default withSQLShellContext(SQLShellPage);