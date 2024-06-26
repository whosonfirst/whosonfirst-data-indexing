# bin

## Configurations

These tools were originally written to expect CLI options. Most of these flags have been removed.

Basically, there are enough different flags and enough details that shouldn't be kept in source control that the CLI flag approach has become a burden. For example all the details used to invoke an ECS task (container name, cluster name, security group, subnet(s), etc.)

Instead, we are using an alternative approach to pull in `.env` files with defaults at runtime. These files are explicitly excluded from source control. For example:

```
WHOAMI=`realpath $0`
FNAME=`basename $WHOAMI`

# Pull in defaults from .env file
if [ -f ${BIN}/${FNAME}.env ]
then
    source ${BIN}/${FNAME}.env
fi
```

Because these tools are cascading (as in one invokes another) the convention is to use the `VARIABLE="${VARIABLE:=default_value}"` syntax. For example:

```
ECS_CONTAINER="${ECS_CONTAINER:=whosonfirst-data-indexing}"
```

This _should_ help to make the individual tools easier to maintain since it removes the need for a lot of boilerplate code to assign CLI variables to the next process in the chain. That's the thinking anyway. In practice things might still change.
