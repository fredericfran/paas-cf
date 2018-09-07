const NotifyClient = require('notifications-node-client').NotifyClient;

const orgs = require('./old_orgs.json');

const NOTIFY_API_KEY = process.env['NOTIFY_API_KEY'];
const NOTIFY_TEMPLATE_ID = process.env['NOTIFY_TEMPLATE_ID'];

function setupNotify() {
  const notify = new NotifyClient(NOTIFY_API_KEY);

  return notify;
}

async function organisationsWithManagers(organisations, remaining) {
  return organisations.reduce((withManagers, current) => {
    if (current.managers > 0) {
      withManagers.push(current);
    } else {
      remaining.push(current);
    }

    return withManagers;
  }, []);
}

async function sendWarningEmail(organisations) {
  const client = setupNotify();

  return organisations.reduce((failing, current) => {
    try {
      const emails = current.org_manager_emails.trim().split(' ');

      emails.forEach(async (email) => {
        await client.sendEmail(NOTIFY_TEMPLATE_ID, email, {
          personalisation: {
            organisation: current.name,
          },
        });
      });
    } catch (err) {
      failing.push(err);
    }
    return failing;
  }, []);
}

async function sendEmails() {
  if (!NOTIFY_API_KEY || !NOTIFY_TEMPLATE_ID) {
    console.error('Both: `NOTIFY_API_KEY` and `NOTIFY_TEMPLATE_ID` must be set whilst using the script.');
    return;
  }

  console.info(`Got ${orgs.length} organisations.\n`);

  const remaining = [];
  const organisations = await organisationsWithManagers(orgs, remaining);

  const failed = await sendWarningEmail(organisations);

  if (failed && failed.length > 0) {
    console.error(`Encountered following errors:`);
    console.error(failed);
  }

  if (remaining.length > 0) {
    console.info('\n', 'You may want to run the following commands manually.', '\n');
    for (const org of remaining) {
      console.log(`cf delete-org ${org.name} -f`);
    }
  }
};

(async () => {
  await sendEmails();
})();
