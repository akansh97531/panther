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

import useRouter from 'Hooks/useRouter';
import { AlertSummary } from 'Generated/schema';
import { Alert, Box, Table } from 'pouncejs';
import RuleDetailsInfo from 'Pages/RuleDetails/RuleDetailsInfo';
import Panel from 'Components/Panel';
import urls from 'Source/urls';
import { extractErrorMessage } from 'Helpers/utils';
import ErrorBoundary from 'Components/ErrorBoundary';
import columns from './columns';
import RuleDetailsPageSkeleton from './Skeleton';
import { useRuleDetails } from './graphql/ruleDetails.generated';

const RuleDetailsPage = () => {
  const { match } = useRouter<{ id: string }>();
  const { history } = useRouter();
  const { error, data, loading } = useRuleDetails({
    fetchPolicy: 'cache-and-network',
    variables: {
      ruleDetailsInput: {
        ruleId: match.params.id,
      },
      alertsForRuleInput: {
        ruleId: match.params.id,
      },
    },
  });

  if (loading && !data) {
    return <RuleDetailsPageSkeleton />;
  }

  if (error) {
    return (
      <Alert
        variant="error"
        title="Couldn't load rule"
        description={
          extractErrorMessage(error) ||
          " An unknown error occured and we couldn't load the rule details from the server"
        }
        mb={6}
      />
    );
  }

  return (
    <article>
      <ErrorBoundary>
        <RuleDetailsInfo rule={data.rule} />
      </ErrorBoundary>
      <Box mt={2} mb={6}>
        <Panel size="large" title="Alerts">
          <ErrorBoundary>
            <Table<Partial<AlertSummary>>
              columns={columns}
              getItemKey={alert => alert.alertId}
              items={data.alerts.alertSummaries}
              onSelect={alert => history.push(urls.logAnalysis.alerts.details(alert.alertId))}
            />
          </ErrorBoundary>
        </Panel>
      </Box>
    </article>
  );
};

export default RuleDetailsPage;
