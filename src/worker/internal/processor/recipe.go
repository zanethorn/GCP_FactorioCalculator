
func (p *Processor) HandleMessage(ctx context.Context, msg *models.Recipe) error {
	_, err := p.Firestore.Collection("recipes").
		Doc(msg.Name).
		Set(ctx, msg)
	return err
}
