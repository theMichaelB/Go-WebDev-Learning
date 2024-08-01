# --- Variables ---

# --- General Targets ---
.PHONY: init plan apply destroy fmt output validate

# --- Main Targets ---
init:  # Initialize tofu working directory
	tofu init 

plan: init  # Create an execution plan
	tofu plan -out=tfplan

apply: plan  # Apply the changes
	tofu apply tfplan 

destroy: init  # Destroy the tofu-managed infrastructure
	tofu destroy -auto-approve

fmt:  # Format tofu configuration files
	tofu fmt -recursive .

validate:  # Validate the tofu configuration files
	tofu validate

output: init  # Display the tofu outputs
	tofu output

# --- Helper Targets ---
clean:  # Remove tofu plan and other generated files
	rm -f tfplan

show: init  # Show the tofu plan
	tofu show tfplan

