/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import React from 'react';
import { LogIntegrationDetails } from 'Source/graphql/fragments/LogIntegrationDetails.generated';
import TablePlaceholder from 'Components/TablePlaceholder';
import { Alert, Box, Button, Card, Flex, Icon, Table } from 'pouncejs';
import Panel from 'Components/Panel';
import { Link } from 'react-router-dom';
import urls from 'Source/urls';
import ErrorBoundary from 'Components/ErrorBoundary';
import { extractErrorMessage } from 'Helpers/utils';
import { useListLogSources } from './graphql/listLogSources.generated';
import columns from './columns';
import EmptyDataFallback from './EmptyDataFallback';

const ListLogSources = () => {
  const { loading, error, data } = useListLogSources();

  if (loading && !data) {
    return (
      <Card p={9}>
        <TablePlaceholder />
      </Card>
    );
  }

  if (error) {
    return (
      <Alert
        variant="error"
        title="Couldn't load your sources"
        description={
          extractErrorMessage(error) ||
          'There was an error when performing your request, please contact support@runpanther.io'
        }
      />
    );
  }

  if (!data.listLogIntegrations.length) {
    return <EmptyDataFallback />;
  }

  return (
    <Box mb={6}>
      <Panel
        title="Log Sources"
        size="large"
        actions={
          <Button size="large" variant="primary" is={Link} to={urls.logAnalysis.sources.create()}>
            <Flex alignItems="center">
              <Icon type="add" size="small" mr={1} />
              Add Source
            </Flex>
          </Button>
        }
      >
        <ErrorBoundary>
          <Table<LogIntegrationDetails>
            items={data.listLogIntegrations}
            getItemKey={item => item.integrationId}
            columns={columns}
          />
        </ErrorBoundary>
      </Panel>
    </Box>
  );
};

export default ListLogSources;
